import BrowserOnly from "@docusaurus/BrowserOnly";
import React from "react";
import { AlsoAvailableAsVideo } from "./AlsoAvailableAsVideo";
import { VideoPlayerWithControls } from "./VideoPlayerWithControls";

export const AppleFireworksTutorial = () => {
  return (
    <BrowserOnly>
      {() => (
        <VideoPlayerWithControls
          playbackId="eEsG2GbZXAqNxays7TqDz24FS1OGCyYVuGYjRg2w2IM"
          poster="https://image.mux.com/eEsG2GbZXAqNxays7TqDz24FS1OGCyYVuGYjRg2w2IM/thumbnail.png"
          onError={(error) => {
            console.log(error);
          }}
          onLoaded={() => undefined}
          onSize={() => undefined}
        />
      )}
    </BrowserOnly>
  );
};

export const TransformsTutorial = () => {
  return (
    <BrowserOnly>
      {() => (
        <AlsoAvailableAsVideo
          playbackId="T0015eD012a4pB26i8veyHA88IRTro2BxVGtGCY3Eq81o"
          thumb="https://image.mux.com/T0015eD012a4pB26i8veyHA88IRTro2BxVGtGCY3Eq81o/thumbnail.png"
          minutes={9}
          title="The 5 basic transformations"
        />
      )}
    </BrowserOnly>
  );
};

export const WarpTextTutorial = () => {
  return (
    <BrowserOnly>
      {() => (
        <AlsoAvailableAsVideo
          playbackId="925Vqjb0002WjmpV7l1qsUXjhKM8vABxW01FKyfAMckX01U"
          thumb="https://image.mux.com/925Vqjb0002WjmpV7l1qsUXjhKM8vABxW01FKyfAMckX01U/thumbnail.png"
          minutes={5}
          title="Cool text effects with warpPath()"
        />
      )}
    </BrowserOnly>
  );
};
