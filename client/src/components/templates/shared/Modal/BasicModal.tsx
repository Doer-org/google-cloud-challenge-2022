import { ReactNode } from 'react';
import { Button } from '../../../atoms/text/Button';

type TProps = {
  children: ReactNode;
};

export const BasicModal = ({ children }: TProps) => {
  return (
    <div className="fixed h-full w-full bg-origin_depth opacity-95 flex flex-col justify-center top-0 left-0 z-50">
      <div className="m-auto w-screen">{children}</div>
    </div>
  );
};
