import { ReactNode, useEffect, useRef, useState } from 'react';

type TProps = {
  children: ReactNode;
  className?: string;
};

export const BasicTemplate = ({ children, className }: TProps) => {
  const el = useRef<HTMLInputElement>(null);
  const browseHeight = document.documentElement.clientHeight;
  const elementHeight = Number(el?.current?.getBoundingClientRect().height);
  const [height, setHeight] = useState<string>('');
  // 初回のレンダリングによる判断
  useEffect(() => {
    // globalの定義を使うとうまく行かないので再定義
    const bh = document.documentElement.clientHeight;
    const elh = Number(el?.current?.getBoundingClientRect().height);
    setHeight(bh > elh ? 'h-screen' : '');
  }, []);
  // リサイズされた際の切り替え
  // 画面幅が変わった時のみ走る
  useEffect(() => {
    const onResize = () => {
      // ここも再定義しないとスタイルの切り替えがうまく行かない
      const bh = document.documentElement.clientHeight;
      const elh = Number(el?.current?.getBoundingClientRect().height);
      setHeight(bh > elh ? 'h-screen' : '');
    };
    window.addEventListener('resize', onResize);
    return () => window.removeEventListener('resize', onResize);
  }, [browseHeight, elementHeight]);

  return (
    <main
      className={`bg-origin flex flex-col justify-center ${className} ${height}`}
    >
      <div
        className={`bg-origin border-4 border-white flex md:m-3 m-2 flex-col justify-center rounded-xl ${height}`}
        ref={el}
      >
        {children}
      </div>
    </main>
  );
};
