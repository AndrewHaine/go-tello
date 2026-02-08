import classNames from "classnames";

interface BtnProps {
  className?: string;
  label?: string;
  children?: React.ReactNode;
  onClick: () => void;
}

export const Btn = (props: BtnProps) => {
  const { className, children, label, onClick } = props;

  return (
    <button
      onClick={onClick}
      className={classNames(
        "px-3 py-2 rounded-full font-black flex flex-col justify-center bg-white/10 leading-5 backdrop-blur-[1px] border-t border-l border-r border-white/40 border-solid p-0 text-lg hover:bg-white/35 active:bg-white/50 cursor-pointer transition-colors",
        className,
      )}
    >
      {children}
      {label ?? null}
    </button>
  );
};
