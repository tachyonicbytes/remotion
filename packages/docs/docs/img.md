---
image: /generated/articles-docs-img.png
title: <Img>
id: img
crumb: "API"
---

The `<Img />` tag can be used like a regular `<img>` HTML tag.

If you use `<Img>`, Remotion will ensure that the image is loaded before rendering the frame. That way you can avoid flickers if the image does not load immediately during rendering.

## API

### `src`

[Put an image into the `public/` folder](/docs/assets) and use [`staticFile()`](/docs/staticfile) to reference it.

```tsx twoslash
import { AbsoluteFill, Img, staticFile } from "remotion";

export const MyComp: React.FC = () => {
  return (
    <AbsoluteFill>
      <Img src={staticFile("hi.png")} />
    </AbsoluteFill>
  );
};
```

You can also load a remote image:

```tsx twoslash
import { AbsoluteFill, Img } from "remotion";

export const MyComp: React.FC = () => {
  return (
    <AbsoluteFill>
      <Img src={"https://picsum.photos/200/300"} />
    </AbsoluteFill>
  );
};
```

### `onError`

To use the `<Img>` tag in a resilient way, handle the error that occurs when an image can not be loaded:

```tsx twoslash
import { AbsoluteFill, Img, staticFile } from "remotion";

export const MyComp: React.FC = () => {
  return (
    <AbsoluteFill>
      <Img
        src={staticFile("hi.png")}
        onError={(event) => {
          // Handle image loading error here
        }}
      />
    </AbsoluteFill>
  );
};
```

If an error occurs, the component must be unmounted or the `src` must be replaced, otherwise the render will time out.

From `v3.3.82`, the image load will first be retried before `onError` is thrown.

## `maxRetries` <AvailableFrom v="v3.3.82"/>

If an image fails to load, it will be retried from `v3.3.82`. The default value is `2`.  
An exponential backoff is being used, with 1000ms delay between the first and second attempt, then 2000ms, then 4000ms and so on.

## Other props

Remotion inherits the props of the regular `<img>` tag, like for example `style`.

## GIFs

Don't use the `<Img>` tag for GIFs, use [`@remotion/gif`](/docs/gif) instead.

## Restrictions

- The maximum resolution that Chrome can display is `2^29` pixels (539 megapixels) <sup><a href="https://stackoverflow.com/questions/57223559/what-is-the-maximum-image-dimensions-supported-in-desktop-chrome#:~:text=than%202%5E29-,(539MP)">[source]</a></sup>. Remotion inherits this restriction.

## See also

- [Source code for this component](https://github.com/remotion-dev/remotion/blob/main/packages/core/src/Img.tsx)
- [Use `<Img>` and `<IFrame>` tags](/docs/use-img-and-iframe)
