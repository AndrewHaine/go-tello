import classNames from "classnames";

interface CircleBtnProps {
  isActive?: boolean;
  className?: string;
  label?: string;
  caption?: string;
  children?: React.ReactNode;
  onClick: () => void;
}

export const CircleBtn = (props: CircleBtnProps) => {
  const { isActive, className, children, caption, label, onClick } = props;

  return (
    <button
      onClick={onClick}
      className={classNames(
        "aspect-square rounded-full font-black flex flex-col justify-center bg-white/20 leading-5 backdrop-blur-[2px] border-t border-l border-r border-white/40 border-solid p-0 text-2xl hover:bg-white/35 active:bg-white/50 cursor-pointer transition-colors",
        className,
        { "bg-white/50": isActive },
      )}
    >
      {children}
      {label ?? null}
      {caption ? <span className="text-xs font-light">{caption}</span> : null}
    </button>
  );
};
