import useDrone from "./hooks/use-drone";
import { Keyboard } from "./components/Keyboard";
import { Telemetry } from "./components/Telemetry";
import { Messages } from "./components/Messages";
import { useRef } from "react";
import { Video } from "./components/Video";
import { Screenshots } from "./components/Screenshots";

const BACKEND_HOST = "http://localhost:8080/api";
const BACKEND_WS_HOST = "ws://localhost:8080/api/ws";

function App() {
  const videoRef = useRef<HTMLVideoElement>(null);

  const { sendCommand, telemetry, messages } = useDrone({
    webSocketHost: BACKEND_WS_HOST,
    videoRef,
  });

  return (
    <>
      <div className="w-full px-6 max-w-300 mx-auto pt-10 grid grid-cols-12 gap-4">
        <h1 className=" text-4xl text-center uppercase font-black col-start-1 col-end-13 mb-2">
          Go Tello Command Centre
        </h1>
        <div id="video" className="col-start-1 col-end-13 md:col-end-8">
          <Video
            batteryLevel={telemetry?.battery}
            host={BACKEND_HOST}
            ref={videoRef}
          />
        </div>
        <div
          id="keyboard"
          className="col-start-1 col-end-13 md:col-end-8 row-start-3"
        >
          <Keyboard onCommand={(command) => sendCommand(command)} />
        </div>
        <div className="flex flex-col col-start-1 md:col-start-8 col-end-13 row-start-2">
          <Telemetry telemetry={telemetry} />
          <Messages messages={messages} />
        </div>
        <div className="col-start-1 md:col-start-8 col-end-13 row-start-3 flex flex-col">
          <Screenshots baseUrl={BACKEND_HOST} />
        </div>
      </div>
    </>
  );
}

export default App;
