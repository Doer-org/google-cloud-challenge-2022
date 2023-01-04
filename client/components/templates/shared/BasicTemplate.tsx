import { ReactNode } from 'react';

type TProps = {
  children: ReactNode;
  className?: string;
};

export const BasicTemplate = ({ children, className }: TProps) => {
  return (
    <main className={`bg-origin h-screen flex flex-col ${className}`}>
      <div className="border-4 border-white flex m-3 flex-col h-screen justify-center rounded-xl">
        {children}
      </div>
    </main>
  );
};
