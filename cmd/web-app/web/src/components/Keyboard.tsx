import { useEffect, useState } from "react";
import { Btn } from "./Btn";
import { CircleBtn } from "./CircleBtn";

interface KeyboardProps {
  onCommand: (command: string) => void;
}

type KeyBindingCategory = "WASD" | "ARROWS";

interface KeyBinding {
  category: KeyBindingCategory;
  label: string;
  caption?: string;
  command: string;
  keyCode: string;
  className: string;
}

const KEY_BINDINGS: Array<KeyBinding> = [
  {
    category: "WASD",
    label: "w",
    keyCode: "KeyW",
    caption: "fwd",
    command: "forward 60",
    className: "col-start-2",
  },
  {
    category: "WASD",
    label: "a",
    keyCode: "KeyA",
    caption: "lft",
    command: "left 60",
    className: "col-start-1 row-start-2",
  },
  {
    category: "WASD",
    label: "s",
    keyCode: "KeyS",
    caption: "bck",
    command: "back 60",
    className: "col-start-2 row-start-2",
  },
  {
    category: "WASD",
    label: "d",
    keyCode: "KeyD",
    caption: "rgt",
    command: "right 60",
    className: "col-start-3 row-start-2",
  },
  {
    category: "ARROWS",
    label: "↑",
    keyCode: "ArrowUp",
    caption: "up",
    command: "up 60 60",
    className: "col-start-2",
  },
  {
    category: "ARROWS",
    label: "←",
    keyCode: "ArrowLeft",
    caption: "ccw",
    command: "ccw 90",
    className: "col-start-1 row-start-2",
  },
  {
    category: "ARROWS",
    label: "↓",
    keyCode: "ArrowDown",
    caption: "dwn",
    command: "down 60",
    className: "col-start-2 row-start-2",
  },
  {
    category: "ARROWS",
    label: "→",
    keyCode: "ArrowRight",
    caption: "cw",
    command: "cw 90",
    className: "col-start-3 row-start-2",
  },
];

export const Keyboard = (props: KeyboardProps) => {
  const { onCommand } = props;
  const [activeKey, setActiveKey] = useState<string | null>(null);

  useEffect(() => {
    window.addEventListener("keydown", (event) => {
      const keybinding = KEY_BINDINGS.find(
        (keybinding) => keybinding.keyCode === event.code,
      );

      if (!keybinding) {
        return;
      }

      setActiveKey(keybinding.keyCode);
      onCommand(keybinding.command);
    });

    window.addEventListener("keyup", () => {
      setActiveKey(null);
    });
  }, [onCommand]);

  return (
    <div className="px-6 py-5 liquid-bg rounded-4xl backdrop-blur-xs grid grid-cols-3 gap-2">
      <div className="grid grid-cols-3 grid-rows-2 gap-2">
        {KEY_BINDINGS.filter(
          (keybinding) => keybinding.category === "WASD",
        ).map((keybinding) => (
          <CircleBtn
            key={keybinding.keyCode}
            isActive={activeKey === keybinding.keyCode}
            className={keybinding.className}
            label={keybinding.label}
            caption={keybinding.caption}
            onClick={() => onCommand(keybinding.command)}
          />
        ))}
      </div>
      <div className="grid grid-cols-3 grid-rows-2 gap-2">
        {KEY_BINDINGS.filter(
          (keybinding) => keybinding.category === "ARROWS",
        ).map((keybinding) => (
          <CircleBtn
            key={keybinding.keyCode}
            isActive={activeKey === keybinding.keyCode}
            className={keybinding.className}
            label={keybinding.label}
            caption={keybinding.caption}
            onClick={() => onCommand(keybinding.command)}
          />
        ))}
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
