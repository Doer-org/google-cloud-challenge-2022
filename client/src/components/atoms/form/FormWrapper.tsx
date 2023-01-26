import { ReactNode } from 'react';

type TProps = {
  children: ReactNode;
  onSubmit: () => void;
};

export const FormWrapper = ({ children, onSubmit }: TProps) => {
  return (
    <form
      className="text-left md:w-1/3 w-4/5 m-auto my-5"
      onSubmit={(event) => {
        event.preventDefault();
        onSubmit();
      }}
    >
      {children}
    </form>
  );
};
