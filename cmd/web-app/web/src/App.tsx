// import { useEffect } from "react";
import useDrone from "./hooks/use-drone";
// import { ReadyState } from "react-use-websocket";
import { Video } from "./components/Video";
import { Keyboard } from "./components/Keyboard";
import { Telemetry } from "./components/Telemetry";
import { Messages } from "./components/Messages";

const BACKEND_HOST = "ws://localhost:8080/ws";

function App() {
  const { sendCommand, telemetry, messages } = useDrone({
    webSocketHost: BACKEND_HOST,
  });

  return (
    <>
      <div className="w-full px-6 max-w-300 mx-auto pt-10 grid grid-cols-12 gap-4">
        <h1 className=" text-4xl text-center uppercase font-black col-start-1 col-end-13 mb-2">
          Go Tello Command Centre
        </h1>
        <div id="video" className="col-start-1 col-end-13 md:col-end-8">
          <Video batteryLevel={telemetry.current?.battery} />
        </div>
        <div
          id="keyboard"
          className="col-start-1 col-end-13 md:col-end-8 row-start-3"
        >
          <Keyboard onCommand={(command) => sendCommand(command)} />
        </div>
        <div className="flex flex-col col-start-1 md:col-start-8 col-end-13 row-start-2">
          <Telemetry telemetry={telemetry.current} />
          <Messages messages={messages.current} />
        </div>
      </div>
    </>
  );
}

export default App;
