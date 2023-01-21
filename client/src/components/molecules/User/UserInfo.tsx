import { useState } from 'react';
import { TypoWrapper } from '../../atoms/text/TypoWrapper';

type TProps = {
  name: string;
  comment?: string;
  image?: string;
  participate?: boolean;
};

export const UserInfo = ({ name, comment, image, participate }: TProps) => {
  return (
    <div className={`${participate ? 'w-20' : 'w-32'}`}>
      {participate ? (
        <>
          {comment ? (
            <div className="relative py-4">
              <div className="bg-white text-black rounded-md comment p-1">
                <TypoWrapper size="small" line="shin">
                  <p className="overflow-x-scroll break-keep whitespace-nowrap">
                    {comment}
                  </p>
                </TypoWrapper>
              </div>
              <span className="before:absolute before:rotate-[180deg] before:-bottom-1 before:left-[25%] before:translate-x-1/2 before:content-[''] before:border-t-[10%] before:border-b-white before:border-[10px] before:border-transparent"></span>
            </div>
          ) : (
            <>
              <div className="bg-white text-black py-2 px-1 rounded-md">
                ...
              </div>
            </>
          )}
        </>
      ) : (
        <></>
      )}
      <div
        className={`${
          participate ? 'w-8 h-8' : 'w-12 h-12'
        } rounded-full bg-orange-500 m-auto my-1`}
      ></div>
      <TypoWrapper size="small" line="shin">
        <p
          className={`overflow-x-scroll w-full break-keep ${
            participate ? '' : 'border-b-2'
          }`}
        >
          {name}
        </p>
      </TypoWrapper>
    </div>
  );
};
