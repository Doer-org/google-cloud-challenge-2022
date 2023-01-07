import { ReactNode } from 'react';
type TProps = {
  children: ReactNode;
};
export const EventWrapper = ({ children }: TProps) => {
  return (
    <div className="m-5 border-accent_border border-8 shadow-2xl bg-origin_depth rounded-md">
      <div className="lg:m-3 m-1 lg:p-3 p-1 border-white border rounded-md">
        {children}
      </div>
    </div>
  );
};
