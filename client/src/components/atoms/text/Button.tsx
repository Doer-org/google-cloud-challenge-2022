import { ReactNode } from 'react';

type TProps = {
  children: ReactNode;
  className?: string;
  onClick?: () => void;
  border?: boolean;
};
export const Button = ({ children, className, onClick, border }: TProps) => {
  return (
    <button onClick={onClick} className={`${className}`}>
      <span className={`${border ? 'border-b-2' : ''} border-white px-3 py-1`}>
        {children}
      </span>
    </button>
  );
};
