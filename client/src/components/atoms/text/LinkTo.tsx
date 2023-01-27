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
export const LinkTo = ({ className, children, href, borderNone, imgPath }: TProps) => {
  return (
    <Link href={`${href}`} className={className}>
      <div className='px-3 py-1'>
        <div className={`${borderNone ? '' : 'border-b-2 border-white'}`}>
          <div className='flex justify-center'>
          {
            imgPath ? <>
              <Image
                src={imgPath}
                alt={'top'}
                width={'25'}
                height={'25'}
                className="m-1"
            />
            </> : <></>
          }
          {children}
          </div>
        </div>
      </div>
    </Link>
  );
};
