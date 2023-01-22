import { ReactNode, useEffect, useRef, useState } from 'react';

type TProps = {
  children: ReactNode;
  className?: string;
};

export const BasicTemplate = ({ children, className }: TProps) => {
  // TODO: globalstateで処理成功時や処理失敗時にUI上にメッセージを表示
  const el = useRef<HTMLInputElement>(null);
  const [browseHeight, setbrowseHeight] = useState(0);
  const [elementHeight, setElementHeight] = useState(0);
  const [height, setHeight] = useState<string>('');

  useEffect(() => {
    const bh = document.documentElement.clientHeight;
    const elh = Number(el.current?.getBoundingClientRect().height);
    setbrowseHeight(bh);
    setElementHeight(elh);
    setHeight(bh > elh ? 'h-screen' : '');
    setTimeout(() => {
      const bh = document.documentElement.clientHeight;
      const elh = Number(el.current?.getBoundingClientRect().height);
      setHeight(bh > elh ? 'h-screen' : '');
    }, 100);
  }, [children, browseHeight, elementHeight]);

  // リサイズされた際の切り替え
  useEffect(() => {
    const onResize = () => {
      const bh = document.documentElement.clientHeight;
      setbrowseHeight(bh);
      const elh = Number(el?.current?.getBoundingClientRect().height);
      setElementHeight(elh);
      setHeight(bh > elh ? 'h-screen' : '');
    };
    window.addEventListener('resize', onResize);
    return () => window.removeEventListener('resize', onResize);
  }, [browseHeight]);

  return (
    <main
      className={`bg-origin flex flex-col justify-center py-2 ${className} ${height}`}
    >
      <div
        className={`bg-origin border-4 border-white flex md:m-3 m-1 flex-col justify-center rounded-xl ${height}`}
      >
        <div ref={el}>{children}</div>
      </div>
    </main>
  );
};
