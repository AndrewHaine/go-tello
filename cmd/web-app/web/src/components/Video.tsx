import { CircleBtn } from "./CircleBtn";

interface VideoProps {
  batteryLevel?: string;
}

export const Video = (props: VideoProps) => {
  const { batteryLevel } = props;

  return (
    <div className="relative aspect-5/3 overflow-hidden rounded-4xl">
      <div className="absolute top-4 right-4 z-10  px-4 py-1 rounded-full flex items-center liquid-bg">
        <span className="bg-red-600 w-3 h-3 rounded-full border-red-400 border-solid border mr-2"></span>
        <span className="font-black">Live</span>
      </div>
      {batteryLevel ? (
        <div className="absolute top-4 left-4 z-10 px-4 py-1 flex items-center rounded-full liquid-bg">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 640 640"
            className="w-6 mr-2"
          >
            <path
              d="M528 192c8.8 0 16 7.2 16 16v224c0 8.8-7.2 16-16 16H112c-8.8 0-16-7.2-16-16V208c0-8.8 7.2-16 16-16zm-416-64c-44.2 0-80 35.8-80 80v224c0 44.2 35.8 80 80 80h416c44.2 0 80-35.8 80-80v-48c17.7 0 32-14.3 32-32v-64c0-17.7-14.3-32-32-32v-48c0-44.2-35.8-80-80-80zm56 112c-13.3 0-24 10.7-24 24v112c0 13.3 10.7 24 24 24h304c13.3 0 24-10.7 24-24V264c0-13.3-10.7-24-24-24z"
              fill="white"
            />
          </svg>
          <span className="font-black">{batteryLevel}%</span>
        </div>
      ) : null}
      <div className="absolute bottom-5 left-[50%] -translate-x-[50%] z-10">
        <CircleBtn caption="snap" onClick={() => console.log("Screenshot")}>
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 640 640"
            className="w-10 mx-4 -mb-1"
            fill="white"
          >
            <path d="M257.1 96C238.4 96 220.9 105.4 210.5 120.9L184.5 160L128 160C92.7 160 64 188.7 64 224L64 480C64 515.3 92.7 544 128 544L512 544C547.3 544 576 515.3 576 480L576 224C576 188.7 547.3 160 512 160L455.5 160L429.5 120.9C419.1 105.4 401.6 96 382.9 96L257.1 96zM250.4 147.6C251.9 145.4 254.4 144 257.1 144L382.8 144C385.5 144 388 145.3 389.5 147.6L422.7 197.4C427.2 204.1 434.6 208.1 442.7 208.1L512 208.1C520.8 208.1 528 215.3 528 224.1L528 480.1C528 488.9 520.8 496.1 512 496.1L128 496C119.2 496 112 488.8 112 480L112 224C112 215.2 119.2 208 128 208L197.3 208C205.3 208 212.8 204 217.3 197.3L250.5 147.5zM320 448C381.9 448 432 397.9 432 336C432 274.1 381.9 224 320 224C258.1 224 208 274.1 208 336C208 397.9 258.1 448 320 448zM256 336C256 300.7 284.7 272 320 272C355.3 272 384 300.7 384 336C384 371.3 355.3 400 320 400C284.7 400 256 371.3 256 336z" />
          </svg>
        </CircleBtn>
      </div>
      <div className="top-0 left-0 object-cover w-full h-full absolute grid grid-cols-1 grid-rows-1">
        <img
          src="/images/static.gif"
          alt="Static image"
          className="w-full h-full object-cover row-start-1 col-start-1"
        />
        <video
          id="cam"
          autoPlay
          className="w-full h-full object-cover top-0 left-0 row-start-1 col-start-1"
          loop
          playsInline
          muted
        >
          <source src="https://www.pexels.com/download/video/5579142/" />
        </video>
      </div>
    </div>
  );
};
