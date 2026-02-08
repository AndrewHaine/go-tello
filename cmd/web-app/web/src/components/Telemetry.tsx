import type { Telemetry as TelloTelemetry } from "../hooks/use-drone";

interface TelemetryProps {
  telemetry?: TelloTelemetry | null;
}

export const Telemetry = (props: TelemetryProps) => {
  const { telemetry } = props;

  return (
    <div className="px-6 py-5 liquid-bg rounded-4xl grid grid-cols-2">
      <strong className="pl-2 border-y border-white/20 mb-1 border-solid">
        Battery
      </strong>
      <div className="border-y border-white/20 mb-1 border-solid text-right">
        {telemetry?.battery ? `${telemetry.battery}%` : "--"}
      </div>
      <strong className="pl-2 border-b border-white/20 mb-1 border-solid">
        Pitch
      </strong>
      <div className="border-b border-white/20 mb-1 border-solid text-right">
        {telemetry?.pitch ? `${telemetry.pitch}°` : "--"}
      </div>
      <strong className="pl-2 border-b border-white/20 mb-1 border-solid">
        Roll
      </strong>
      <div className="border-b border-white/20 mb-1 border-solid text-right">
        {telemetry?.roll ? `${telemetry.roll}°` : "--"}
      </div>
      <strong className="pl-2 border-b border-white/20 mb-1 border-solid">
        Yaw
      </strong>
      <div className="border-b border-white/20 mb-1 border-solid text-right">
        {telemetry?.yaw ? `${telemetry.yaw}°` : "--"}
      </div>
      <strong className="pl-2 border-b border-white/20 mb-1 border-solid">
        Temp
      </strong>
      <div className="border-b border-white/20 mb-1 border-solid text-right">
        {telemetry?.temp_low && telemetry.temp_high
          ? `${telemetry.temp_low}°C - ${telemetry.temp_high}°C`
          : "--"}
      </div>
      <strong className="pl-2 border-b border-white/20 border-solid">
        Height
      </strong>
      <div className="border-b border-white/20 border-solid text-right">
        {telemetry?.height ? `${telemetry.height}cm` : "--"}
      </div>
    </div>
  );
};
