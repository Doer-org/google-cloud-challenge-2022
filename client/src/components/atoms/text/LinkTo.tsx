import Link from 'next/link';
import { ReactNode } from 'react';
import Image from 'next/image';
type TProps = {
  className?: string;
  children: ReactNode;
  href: string;
  borderNone?: boolean;
  imgPath?: string;
};
export const LinkTo = ({
  className,
  children,
  href,
  borderNone,
  imgPath,
}: TProps) => {
  return (
    <Link href={`${href}`} className={`${className} inline-block`}>
      <div className="py-1">
        <div className={`${borderNone ? '' : 'border-b-2 border-white'}`}>
          <div className="py-1 flex justify-center items-center px-2">
            {imgPath ? (
              <>
                <Image
                  src={imgPath}
                  alt={'top'}
                  width={'25'}
                  height={'25'}
                  className="p-1 pt-2"
                />
              </>
            ) : (
              <></>
            )}
            {children}
          </div>
        </div>
      </div>
    </Link>
  );
};
