import { TypoWrapper } from '../../atoms/text/TypoWrapper';
type TProps = {
  eventName: string;
  detail: string;
  limitTime: string;
};
export const EventBasicInfo = ({ eventName, detail, limitTime }: TProps) => {
  const formatDateString = (dateString: string) => {
    return dateString.replace(
      /(\d{4})-(\d{2})-(\d{2})T(\d{2}):(\d{2}):\d{2}\+\d{2}:\d{2}/,
      '$1年$2月$3日$4時$5分'
    );
  };
  return (
    <>
      <div className="lg:mx-auto lg:w-2/3">
        <TypoWrapper line="bold" size="so-large">
          <h1 className="my-5">{eventName}</h1>
        </TypoWrapper>
        {limitTime ? (
          <TypoWrapper line="shin" size="small">
            <p className="mb-2">
              {`${'（'}`}
              {formatDateString(limitTime) + 'まで'}
              {`${'）'}`}
            </p>
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
