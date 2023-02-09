import { useEffect, useState } from 'react';
import { TypoWrapper } from '../../atoms/text/TypoWrapper';
type TProps = {
  eventName: string;
  detail: string;
  limitTime: string;
};
export const EventBasicInfo = ({ eventName, detail, limitTime }: TProps) => {
  const [time, setTime] = useState('');
  useEffect(() => {
    setTime(`(${new Date(limitTime).toLocaleString()} まで)`);
  }, [limitTime]);
  return (
    <>
      <div className="lg:mx-auto lg:w-2/3">
        <TypoWrapper line="bold" size="so-large">
          <h1 className="my-5">{eventName}</h1>
        </TypoWrapper>
        {limitTime ? (
          <TypoWrapper line="shin" size="small">
            <p className="mb-2">{time}</p>
          </TypoWrapper>
        ) : (
          <></>
        )}
        <TypoWrapper size="small" line="shin">
          <p className="text-center lg:mx-10 mx-2 mt-2 mb-10">{detail}</p>
        </TypoWrapper>
      </div>
    </>
  );
};
