import { ReactNode } from 'react';

type TProps = {
  children: ReactNode;
  className?: string;
};

export const BasicTemplate = ({ children, className }: TProps) => {
  return (
    <main
      className={`bg-origin flex flex-col h-screen justify-center ${className}`}
    >
      {children}
    </main>
  );
};
