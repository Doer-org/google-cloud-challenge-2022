import { ReactNode } from 'react';

type TProps = {
  children: ReactNode;
  className?: string;
  onClick?: () => void;
  disable?: boolean;
  border?: boolean;
};
export const Button = ({
  children,
  className,
  onClick,
  border,
  disable,
}: TProps) => {
  return (
    <button onClick={onClick} className={`${className}`} disabled={disable}>
      <span className={`${border ? 'border-b-2' : ''} border-white px-3 py-1`}>
        {children}
      </span>
    </button>
  );
};
