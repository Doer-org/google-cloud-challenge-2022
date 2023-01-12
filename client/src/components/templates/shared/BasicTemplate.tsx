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
  // リサイズされた際の切り替え
  // 画面幅が変わった時のみ走る
  useEffect(() => {
    const bh = document.documentElement.clientHeight;
    const elh = Number(el?.current?.getBoundingClientRect().height);
    setHeight(bh > elh ? 'h-screen' : '');
    // TODO:高さが800前後の時に初回レンダリングでうまく高さが判定できてない
    // 高さが本来の高さよりも低くなっていることからh-screenが適用されてしまっている
    // iPad Airとかだと大丈夫
    window.onload = () => {
      const bh = document.documentElement.clientHeight;
      const elh = Number(el?.current?.getBoundingClientRect().height);
      setHeight(bh > elh ? 'h-screen' : '');
    };
  }, []);
  useEffect(() => {
    const bh = document.documentElement.clientHeight;
    const elh = Number(el?.current?.getBoundingClientRect().height);
    setbrowseHeight(bh);
    setElementHeight(elh);
    setHeight(browseHeight > elementHeight ? 'h-screen' : '');
    const onResize = () => {
      // ここも再定義しないとスタイルの切り替えがうまく行かない
      const bh = document.documentElement.clientHeight;
      const elh = Number(el?.current?.getBoundingClientRect().height);
      setHeight(bh > elh ? 'h-screen' : '');
    };
    window.addEventListener('resize', onResize);
    return () => window.removeEventListener('resize', onResize);
  }, [browseHeight, elementHeight, height]);

  return (
    <main
      className={`bg-origin flex flex-col justify-center py-2 ${className} ${height}`}
    >
      <div
        className={`bg-origin border-4 border-white flex md:m-3 m-1 flex-col justify-center rounded-xl ${height}`}
        ref={el}
      >
        {children}
      </div>
    </main>
  );
};
