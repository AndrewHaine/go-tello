import { useQueryClient } from "@tanstack/react-query";
import {
  useCallback,
  useEffect,
  useRef,
  useState,
  type RefObject,
} from "react";
import { SCREENSHOTS_QUERY_KEY } from "../components/Screenshots";

interface UseDroneCommandsOptions {
  webSocketHost: string;
  videoRef: RefObject<HTMLVideoElement | null>;
}

type EventType =
  | "log.created"
  | "command.requested"
  | "connection.updated"
  | "telemetry.updated"
  | "video.offer"
  | "video.answer"
  | "video.ice-candidate"
  | "screenshot.added";

interface Event {
  event: EventType;
  timestamp: string;
}

export interface Telemetry {
  battery: string;
  pitch: string;
  roll: string;
  yaw: string;
  temp_high: string;
  temp_low: string;
  height: string;
}

interface TelemetryEvent extends Event {
  payload: Telemetry;
}

const isTelemetryEvent = (event: Event): event is TelemetryEvent => {
  return event.event === "telemetry.updated";
};

export interface Message {
  message: string;
  time: string;
}

interface MessageEvent extends Event {
  payload: Message;
}

const isMessageEvent = (event: Event): event is MessageEvent => {
  return event.event === "log.created";
};

interface VideoPeerConnectionOfferEvent extends Event {
  payload: RTCSessionDescriptionInit;
}

const isVideoPeerConnectionOfferEvent = (
  event: Event,
): event is VideoPeerConnectionOfferEvent => {
  return event.event === "video.offer";
};

interface VideoPeerConnectionIceCandidateEvent extends Event {
  payload: RTCIceCandidateInit;
}

const isVideoPeerConnectionIceCandidateEvent = (
  event: Event,
): event is VideoPeerConnectionIceCandidateEvent => {
  return event.event === "video.ice-candidate";
};

export default function useDrone(options: UseDroneCommandsOptions) {
  const { videoRef, webSocketHost } = options;
  const [telemetry, setTelemetry] = useState<Telemetry | null>(null);
  const [messages, setMessages] = useState<Array<Message>>([]);
  const [readyState, setReadyState] = useState<number>(WebSocket.CONNECTING);
  const queryClient = useQueryClient();

  const wsRef = useRef<WebSocket | null>(null);
  const videoPeerConnection = useRef<RTCPeerConnection | null>(null);

  const handleVideoOffer = async (event: VideoPeerConnectionOfferEvent) => {
    const pc = videoPeerConnection.current;
    if (!pc) return;

    await pc.setRemoteDescription(event.payload);
    const answer = await pc.createAnswer();
    await pc.setLocalDescription(answer);

    wsRef.current?.send(
      JSON.stringify({
        event: "video.answer",
        payload: answer,
      }),
    );
  };

  const handleIceCandidate = async (
    event: VideoPeerConnectionIceCandidateEvent,
  ) => {
    if (videoPeerConnection.current) {
      await videoPeerConnection.current.addIceCandidate(event.payload);
    }
  };

  useEffect(() => {
    const ws = new WebSocket(webSocketHost);
    wsRef.current = ws;

    ws.onopen = () => setReadyState(WebSocket.OPEN);
    ws.onclose = () => setReadyState(WebSocket.CLOSED);
    ws.onerror = () => setReadyState(WebSocket.CLOSED);

    ws.onmessage = (event) => {
      const parsedEvent = JSON.parse(event.data) as Event;

      if (isTelemetryEvent(parsedEvent)) {
        setTelemetry(parsedEvent.payload);
        return;
      }

      if (isMessageEvent(parsedEvent)) {
        setMessages((prev) => [...prev, parsedEvent.payload]);
        return;
      }

      if (isVideoPeerConnectionOfferEvent(parsedEvent)) {
        handleVideoOffer(parsedEvent);
        return;
      }

      if (isVideoPeerConnectionIceCandidateEvent(parsedEvent)) {
        handleIceCandidate(parsedEvent);
      }

      if (parsedEvent.event === "screenshot.added") {
        queryClient.invalidateQueries({
          queryKey: [SCREENSHOTS_QUERY_KEY],
        });
      }
    };

    const pc = new RTCPeerConnection();
    videoPeerConnection.current = pc;

    pc.ontrack = (event) => {
      if (videoRef.current) {
        videoRef.current.srcObject = event.streams[0];
        videoRef.current.play();
      }
    };

    pc.onicecandidate = (event) => {
      ws.send(
        JSON.stringify({
          event: "video.ice-candidate",
          payload: event.candidate,
        }),
      );
    };

    return () => {
      ws.close();
      pc.close();
    };
    // eslint-disable-next-line
  }, [webSocketHost, videoRef]);

  const sendCommand = useCallback((command: string) => {
    if (wsRef.current?.readyState === WebSocket.OPEN) {
      wsRef.current.send(
        JSON.stringify({
          event: "command.requested",
          payload: { command },
          timestamp: new Date().toISOString(),
        }),
      );
    }
  }, []);

  return { readyState, sendCommand, telemetry, messages };
}
