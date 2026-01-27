import { useEffect } from "react";
import useWebSocket from "react-use-websocket";

interface UseDroneCommandsOptions {
  webSocketHost: string;
}

// interface Telemetry {
//   battery: number;
//   pitch: number;
//   roll: number;
//   yaw: number;
//   temp: string;
//   height: number;
// }

// const BLANK_TELEMETRY: Telemetry = {
//   battery: 0,
//   pitch: 0,
//   roll: 0,
//   yaw: 0,
//   temp: "",
//   height: 0,
// };

// interface Message {
//   time: Date;
//   message: string;
// }

// const DroneStatus = {
//   ONLINE: "ONLINE",
//   OFFLINE: "OFFLINE",
// } as const;

export default function useDrone(options: UseDroneCommandsOptions) {
  const { webSocketHost } = options;
  // const [droneStatus, setDroneStatus] = useState<DroneStatus>("OFFLINE");
  // const [telemetry, setTelemetry] = useState(BLANK_TELEMETRY);
  // const [messages, setMessages] = useState<Array<Message>>([]);

  const {
    sendMessage: sendCommand,
    lastMessage,
    readyState,
  } = useWebSocket(webSocketHost);

  useEffect(() => {
    console.log(lastMessage);
  }, [lastMessage]);

  return {
    readyState,
    sendCommand,
  };
}
