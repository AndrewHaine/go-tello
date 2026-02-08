import { useCallback, useEffect, useRef } from "react";
import useWebSocket from "react-use-websocket";

interface UseDroneCommandsOptions {
  webSocketHost: string;
}

type EventType =
  | "log.created"
  | "command.requested"
  | "connection.updated"
  | "telemetry.updated";

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

interface CommandEvent extends Event {
  payload: { command: string };
}

// const DroneStatus = {
//   ONLINE: "ONLINE",
//   OFFLINE: "OFFLINE",
// } as const;

export default function useDrone(options: UseDroneCommandsOptions) {
  const { webSocketHost } = options;
  // const [droneStatus, setDroneStatus] = useState<DroneStatus>("OFFLINE");
  const telemetry = useRef<Telemetry | null>(null);
  const messages = useRef<Array<Message>>([]);

  const { sendMessage, lastMessage, readyState } = useWebSocket(webSocketHost);

  useEffect(() => {
    if (!lastMessage?.data) {
      return;
    }

    const lastEvent = JSON.parse(lastMessage?.data) as Event;

    if (isTelemetryEvent(lastEvent)) {
      telemetry.current = lastEvent.payload;
      return;
    }

    if (isMessageEvent(lastEvent)) {
      messages.current = [...messages.current, lastEvent.payload];
      return;
    }
  }, [lastMessage]);

  const sendCommand = useCallback(
    (command: string) => {
      const event: CommandEvent = {
        event: "command.requested",
        payload: {
          command,
        },
        timestamp: new Date().toISOString(),
      };
      sendMessage(JSON.stringify(event));
    },
    [sendMessage],
  );

  return {
    readyState,
    sendCommand,
    telemetry,
    messages,
  };
}
