import { Btn } from "./Btn";
import { CircleBtn } from "./CircleBtn";

interface KeyboardProps {
  onCommand: (command: string) => void;
}

export const Keyboard = (props: KeyboardProps) => {
  const { onCommand } = props;

  return (
    <div className="px-6 py-5 liquid-bg rounded-4xl backdrop-blur-xs grid grid-cols-3 gap-2">
      <div className="grid grid-cols-3 grid-rows-2 gap-2">
        <CircleBtn
          className="col-start-2"
          label="w"
          caption="fwd"
          onClick={() => onCommand("forward 60")}
        />
        <CircleBtn
          className="col-start-1 row-start-2"
          label="a"
          caption="lft"
          onClick={() => onCommand("left 60")}
        />
        <CircleBtn
          className="col-start-2 row-start-2"
          label="s"
          caption="bck"
          onClick={() => onCommand("back 60")}
        />
        <CircleBtn
          className="col-start-3 row-start-2"
          label="d"
          caption="rgt"
          onClick={() => onCommand("right 60")}
        />
      </div>
      <div className="grid grid-cols-3 grid-rows-2 gap-2">
        <CircleBtn
          className="col-start-2"
          label="↑"
          caption="up"
          onClick={() => onCommand("up 60")}
        />
        <CircleBtn
          className="col-start-1 row-start-2"
          label="←"
          caption="ccw"
          onClick={() => onCommand("ccw 90")}
        />
        <CircleBtn
          className="col-start-2 row-start-2"
          label="↓"
          caption="dwn"
          onClick={() => onCommand("down 60")}
        />
        <CircleBtn
          className="col-start-3 row-start-2"
          label="→"
          caption="cw"
          onClick={() => onCommand("cw 90")}
        />
      </div>
      <div className="flex flex-col justify-center gap-2 pl-2">
        <Btn
          className="w-full"
          label="Take Off"
          onClick={() => onCommand("takeoff")}
        />
        <Btn
          className="w-full"
          label="Land"
          onClick={() => onCommand("land")}
        />
        <Btn
          className="w-full bg-red-500/65! hover:bg-red-500/75! active:bg-red-500! border-red-500!"
          label="EMERGENCY"
          onClick={() => onCommand("emergency")}
        />
      </div>
    </div>
  );
};
