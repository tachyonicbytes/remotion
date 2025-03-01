---
image: /generated/articles-docs-composition.png
id: composition
title: <Composition>
crumb: "API"
---

This is the component to use to register a video to make it renderable and make it show up in the sidebar of the Remotion development interface.

A composition represents the video you want to create, as a collection of clips (for example, several `<Sequence>`) that will play back to back to form your video.

```tsx twoslash title="src/Root.tsx"
const Component: React.FC = () => null;
// ---cut---

import { Composition } from "remotion";

export const RemotionRoot: React.FC = () => {
  return (
    <>
      <Composition
        component={Component}
        durationInFrames={300}
        width={1080}
        height={1080}
        fps={30}
        id="test-render"
        defaultProps={{}}
      />
      {/* Additional compositions can be rendered */}
    </>
  );
};
```

## API

A `<Composition />` should be placed within a fragment of the root component (which is registered using [`registerRoot()`](/docs/register-root)).

The component takes the following props:

### `id`

ID of the composition, as shown in the sidebar and also a unique identifier of the composition that you need to specify if you want to render it. The ID can only contain letters, numbers and `-`.

### `fps`

At how many frames the composition should be rendered.

### `durationInFrames`

How many frames the composition should be long.

### `height`

Height of the composition in pixels.

### `width`

Width of the composition in pixels.

### `component` **or** `lazyComponent`

Pass the component in directly **or** pass a function that returns a dynamic import. Passing neither or both of the props is an error.

:::note
If you use `lazyComponent`, Remotion will use React Suspense to load the component. Components will be compiled by Webpack as they are needed, which will reduce startup time of Remotion. If you use `lazyComponent`, you need to use a default export for your component. This is a restriction of React Suspense.
:::

### `defaultProps`

_optional_

Give your component default props that will be shown in the preview. You can override these props during render using a CLI flag.

:::note
Type your components using the `React.FC<{}>` type and the `defaultProps` prop will be typesafe.
:::

:::note
Passing huge objects to `defaultProps` can be slow. [Learn how to avoid it.](/docs/troubleshooting/defaultprops-too-big)

:::

## Example using `component`

```tsx twoslash
// @allowUmdGlobalAccess
// @filename: ./MyComp.tsx
export const MyComp = () => <></>;

// @filename: index.tsx
// ---cut---
import { Composition } from "remotion";
import { MyComp } from "./MyComp";

export const MyVideo = () => {
  return (
    <>
      <Composition
        id="my-comp"
        component={MyComp}
        width={1080}
        height={1080}
        fps={30}
        durationInFrames={3 * 30}
      />
    </>
  );
};
```

## Example using `lazyComponent`

```tsx
export const MyVideo = () => {
  return (
    <>
      <Composition
        id="my-comp"
        lazyComponent={() => import("./LazyComponent")}
        width={1080}
        height={1080}
        fps={30}
        durationInFrames={3 * 30}
      />
    </>
  );
};
```

## Organize compositions using folders

You can use the [`<Folder />`](/docs/folder) component to organize your compositions in the sidebar.

```tsx twoslash
import React from "react";
const Component: React.FC = () => null;
// ---cut---
import { Composition, Folder } from "remotion";

export const Video = () => {
  return (
    <>
      <Folder name="Visuals">
        <Composition
          id="CompInFolder"
          durationInFrames={100}
          fps={30}
          width={1080}
          height={1080}
          component={Component}
        />
      </Folder>
      <Composition
        id="CompOutsideFolder"
        durationInFrames={100}
        fps={30}
        width={1080}
        height={1080}
        component={Component}
      />
    </>
  );
};
```

## See also

- [Source code for this component](https://github.com/remotion-dev/remotion/blob/main/packages/core/src/Composition.tsx)
- [registerRoot()](/docs/register-root)
- [The fundamentals](/docs/the-fundamentals)
- [CLI options](/docs/cli)
- [`<Sequence />`](/docs/sequence)
- [`<Still />`](/docs/still)
- [`<Folder />`](/docs/folder)
- [Avoid huge payloads for `defaultProps`](/docs/troubleshooting/defaultprops-too-big)
