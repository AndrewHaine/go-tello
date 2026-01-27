import classNames from "classnames";

interface KeyboardButtonProps {
  letter: string;
  label?: string;
  className?: string;
  onClick: (letter: string) => void;
}

export default function KeyboardButton(props: KeyboardButtonProps) {
  const { letter, label, className, onClick } = props;

  return (
    <button
      className={classNames(
        "border-white border-solid border-2 flex flex-col justify-center items-center aspect-square cursor-pointer select-none",
        className,
      )}
      onClick={() => onClick(letter)}
    >
      <strong>{letter}</strong>
      {label}
    </button>
  );
}
