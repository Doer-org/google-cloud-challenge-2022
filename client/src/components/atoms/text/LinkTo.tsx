import Link from 'next/link';
import { ReactNode } from 'react';
type TProps = {
  className?: string;
  children: ReactNode;
  href: string;
  borderNone?: boolean;
};
export const LinkTo = ({ className, children, href, borderNone }: TProps) => {
  return (
    <Link href={`${href}`} className={className}>
      <span
        className={`${borderNone ? '' : 'border-b-2 border-white'} px-3 py-1`}
      >
        {children}
      </span>
    </Link>
  );
};
