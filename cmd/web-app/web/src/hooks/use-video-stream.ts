import type { RefObject } from "react";

export default function useVideoStream(
  videoRef: RefObject<HTMLVideoElement | null>,
) {
  const peerConn = new RTCPeerConnection();

  peerConn.ontrack = (event) => {
    if (!videoRef.current) {
      return;
    }

    const video = videoRef.current;
    video.srcObject = event.streams[0];
    video.play();
  };
}
