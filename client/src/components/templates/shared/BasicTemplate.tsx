import { ReactNode, useEffect, useRef, useState } from 'react';

type TProps = {
  children: ReactNode;
  className?: string;
};

export const BasicTemplate = ({ children, className }: TProps) => {
  return (
    <main
      className={`flex flex-col justify-center h-fit min-h-full box-border py-2 ${className} border-4 border-white rounded-xl box-border`}
    >
      {children}
    </main>
  );
};
