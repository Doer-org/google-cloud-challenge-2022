import React from 'react';
import { TypoWrapper } from '../../atoms/text/TypoWrapper';
type TProps = {
  name: string;
  isParticipate: boolean | undefined;
  full?: boolean;
};
export const UserName = ({ name, isParticipate, full }: TProps) => {
  return (
    <>
      {full ? (
        <TypoWrapper size="so-large" line="bold">
          <p className={`w-full ${isParticipate ? '' : 'border-b-2'}`}>
            {name}
          </p>
        </TypoWrapper>
      ) : (
        <TypoWrapper size="small" line="shin">
          <p
            className={`overflow-x-scroll w-full break-keep ${
              isParticipate ? '' : 'border-b-2'
            }`}
          >
            {name}
          </p>
        </TypoWrapper>
      )}
    </>
  );
};
