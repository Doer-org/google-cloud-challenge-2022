import { ReactNode } from 'react';

type TProps = {
  children: ReactNode;
  className?: string;
  onClick: () => void;
};
export const Button = ({ children, className, onClick }: TProps) => {
  return (
    <button onClick={onClick} className={`${className}`}>
      <span className="border-b-2 border-white px-3 py-1">{children}</span>
    </button>
  );
};
