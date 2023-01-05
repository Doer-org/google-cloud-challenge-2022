import { ReactNode } from 'react';

type TProps = {
  children: ReactNode;
  className?: string;
};

export const EventFormTemplate = ({ children, className }: TProps) => {
  return (
    <main className={`bg-origin flex lg:h-screen flex-col ${className}`}>
      <div className="bg-origin border-4 border-white flex md:m-3 m-2 flex-col lg:h-screen justify-center rounded-xl">
        {children}
      </div>
    </main>
  );
};
