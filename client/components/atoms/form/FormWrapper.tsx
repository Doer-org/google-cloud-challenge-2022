import { ReactNode } from 'react';

type TProps = {
  children: ReactNode;
};

export const FormWrapper = ({ children }: TProps) => {
  return <div className="text-left md:w-1/3 w-4/5 m-auto my-5">{children}</div>;
};
