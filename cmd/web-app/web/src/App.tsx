import { useEffect } from "react";
import KeyboardButton from "./components/KeyboardButton";
import useDrone from "./hooks/use-drone";
import { ReadyState } from "react-use-websocket";

const BACKEND_HOST = "ws://localhost:8080/ws";

function App() {
  const { sendCommand, readyState } = useDrone({ webSocketHost: BACKEND_HOST });

  useEffect(() => {
    console.log(ReadyState[readyState]);
  }, [readyState]);

  const handleButtonClick = (key: string) => {
    switch (key) {
      case "w":
        sendCommand("forward 50");
        break;
      case "a":
        sendCommand("left 50");
        break;
      case "s":
        sendCommand("back 50");
        break;
      case "d":
        sendCommand("right 50");
        break;
      case "↑":
        sendCommand("up 50");
        break;
      case "←":
        sendCommand("ccw 90");
        break;
      case "↓":
        sendCommand("down 50");
        break;
      case "→":
        sendCommand("cw 90");
        break;
    }
  };

  return (
    <>
      <div className="w-full max-w-300 mx-auto pt-12">
        <h1 className="font-sixtyfour text-4xl text-center">
          Go Tello Command Centre
        </h1>
        <div className="mt-8 relative w-full max-w-3xl mx-auto">
          <video
            id="cam"
            className="aspect-video bg-gray-800"
            autoPlay
            playsInline
            muted
          ></video>
          <span className="absolute left-0 top-0 w-[50%] aspect-square z-10"></span>
        </div>
        <div className="grid grid-cols-[300px_1fr_300px] mt-10 px-12">
          <div className="grid gap-1 grid-cols-3 grid-rows-2">
            <KeyboardButton
              className="col-start-2"
              letter="w"
              label="forwards"
              onClick={handleButtonClick}
            />
            <KeyboardButton
              className="col-start-1 row-start-2"
              letter="a"
              label="left"
              onClick={handleButtonClick}
            />
            <KeyboardButton
              className="col-start-2 row-start-2"
              letter="s"
              label="backward"
              onClick={handleButtonClick}
            />
            <KeyboardButton
              className="col-start-3 row-start-2"
              letter="d"
              label="right"
              onClick={handleButtonClick}
            />
          </div>
          <div className="flex flex-col justify-center gap-4 px-10">
            <button
              className="bg-green-500 font-black px-2 py-1 rounded-2xl"
              onClick={() => sendCommand("takeoff")}
            >
              Take Off
            </button>
            <button
              className="bg-amber-500 font-black px-2 py-1 rounded-2xl"
              onClick={() => sendCommand("land")}
            >
              Land
            </button>
            <button
              className="bg-red-500 font-black px-2 py-1 rounded-2xl"
              onClick={() => sendCommand("emergency")}
            >
              Emergency
            </button>
            <button className="bg-blue-500 font-black px-2 py-1 rounded-2xl">
              Toggle Video
            </button>
          </div>
          <div className="grid gap-1 grid-cols-3 grid-rows-2">
            <KeyboardButton
              className="col-start-2"
              letter="↑"
              label="up"
              onClick={handleButtonClick}
            />
            <KeyboardButton
              className="col-start-1 row-start-2"
              letter="←"
              label="yaw left"
              onClick={handleButtonClick}
            />
            <KeyboardButton
              className="col-start-2 row-start-2"
              letter="↓"
              label="down"
              onClick={handleButtonClick}
            />
            <KeyboardButton
              className="col-start-3 row-start-2"
              letter="→"
              label="yaw right"
              onClick={handleButtonClick}
            />
          </div>
        </div>
      </div>
    </>
  );
}

export default App;
