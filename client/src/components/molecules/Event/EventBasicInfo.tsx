import { LinkTo } from '../../atoms/text/LinkTo';
import { TypoWrapper } from '../../atoms/text/TypoWrapper';
type TProps = {
  eventName: string;
  detail: string;
};
export const EventBasicInfo = ({ eventName, detail }: TProps) => {
  return (
    <>
      <div className="lg:mx-auto lg:w-1/3">
        <TypoWrapper line="bold" size="so-large">
          <h1 className="my-5">{eventName}</h1>
        </TypoWrapper>
        <TypoWrapper size="small" line="shin">
          <p className="text-center lg:mx-10 mx-2 mt-2 mb-10">{detail}</p>
        </TypoWrapper>
      </div>
    </>
  );
};
