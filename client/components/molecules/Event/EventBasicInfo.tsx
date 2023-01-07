import { TypoWrapper } from '../../atoms/text/TypoWrapper';
import { UserIcon } from '../User/UserIcon';
type TProps = {
  eventName: string;
  detail: string;
};
export const EventBasicInfo = ({ eventName, detail }: TProps) => {
  return (
    <>
      <div className="lg:mx-auto lg:w-1/3">
        <TypoWrapper line="bold">
          <h1 className="my-5">{eventName}</h1>
        </TypoWrapper>
        <TypoWrapper size="small" line="shin">
          <p className="text-left lg:mx-10 mx-2 my-2">{detail}</p>
        </TypoWrapper>
      </div>
    </>
  );
};
