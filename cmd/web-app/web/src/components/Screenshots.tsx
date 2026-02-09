import { useQuery } from "@tanstack/react-query";
import { format } from "date-fns";
import { useState } from "react";
import { createPortal } from "react-dom";

export const SCREENSHOTS_QUERY_KEY = "screenshots";

type ScreenshotData = { name: string; date: string };
type ScreenshotRes = { screenshots: Array<ScreenshotData> };

interface ScreenshotsProps {
  baseUrl: string;
}

const ScreenshotsWrapper = (props: { children: React.ReactNode }) => (
  <div className="px-6 py-5 liquid-bg flex flex-col rounded-4xl grow">
    <span className="font-black text-xl">Screenshots</span>
    {props.children}
  </div>
);

export const Screenshots = (props: ScreenshotsProps) => {
  const { baseUrl } = props;

  const { isSuccess, isFetching, data } = useQuery<ScreenshotRes>({
    queryKey: [SCREENSHOTS_QUERY_KEY],
    queryFn: () => fetch(`${baseUrl}/screenshots`).then((r) => r.json()),
  });

  const [activeScreenshot, setActiveScreenshot] =
    useState<ScreenshotData | null>(null);

  if ((isSuccess || isFetching) && data && data.screenshots.length) {
    return (
      <ScreenshotsWrapper>
        {activeScreenshot
          ? createPortal(
              <>
                <div
                  className="fixed z-50 top-0 left-0 w-screen h-screen bg-black/15"
                  onClick={() => setActiveScreenshot(null)}
                />
                <div className="fixed z-60 top-10 left-[50%] w-[80%] max-w-150 translate-x-[-50%] bg-white pt-3 px-4 pb-4">
                  <img src={activeScreenshot.name} alt="Screenshot" />
                  <p className="mt-3 text-black font-bold text-center">
                    {format(
                      new Date(activeScreenshot.date),
                      "EEEE do LLL y @ HH:mm:ss",
                    )}
                  </p>
                </div>
              </>,
              document.body,
            )
          : null}
        <div className="mt-1 flex flex-wrap gap-1">
          {data.screenshots.map((screenshot) => (
            <a
              key={screenshot.date}
              onClick={() => setActiveScreenshot(screenshot)}
              className="block w-20 bg-white pt-1.5 px-1 pb-3"
            >
              <img src={screenshot.name} className="w-full" />
            </a>
          ))}
        </div>
      </ScreenshotsWrapper>
    );
  }

  return (
    <ScreenshotsWrapper>
      <p>No Screenshots</p>
    </ScreenshotsWrapper>
  );
};
