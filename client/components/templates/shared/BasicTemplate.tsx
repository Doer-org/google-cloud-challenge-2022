import { ReactNode } from 'react';

type TProps = {
  children: ReactNode;
  className?: string;
};

export const BasicTemplate = ({ children, className }: TProps) => {
  // h-screenすると要素が100vh以内でうまく配置されるが100vhを超えてしまうとうまく配置されない（要素がはみ出て見えなくなる）
  // これを今はレスポンシブで対応しているが、これだと文が長かった時とかに見えない部分が出てくるのかもしれない
  //  viewの子要素のサイズを見て１００vhするかどうかの判定コードを書きたい
  return (
    <main className={`bg-origin flex lg:h-screen flex-col h-full ${className}`}>
      <div className="bg-origin border-4 border-white flex md:m-3 m-2 flex-col lg:h-screen justify-center rounded-xl">
        {children}
      </div>
    </main>
  );
};
