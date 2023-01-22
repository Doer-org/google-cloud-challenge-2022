import React from 'react';
import { TypoWrapper } from '../../atoms/text/TypoWrapper';
type TProps = {
  name: string;
  full?: boolean;
};
export const UserName = ({ name, full }: TProps) => {
  return (
    <>
      {full ? (
        <TypoWrapper size="large" line="bold">
          <p className={`w-full`}>{name}</p>
        </TypoWrapper>
      ) : (
        <TypoWrapper size="small" line="shin">
          <p className={`overflow-x-scroll w-full break-keep`}>{name}</p>
        </TypoWrapper>
      )}
    </>
  );
};
