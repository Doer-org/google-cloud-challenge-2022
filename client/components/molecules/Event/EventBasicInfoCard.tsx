import { LinkTo } from '../../atoms/text/LinkTo';
import { TypoWrapper } from '../../atoms/text/TypoWrapper';
type TProps = {
  id: string;
  eventName: string;
  detail: string;
};
export const EventBasicInfoCard = ({ id, eventName, detail }: TProps) => {
  return (
    <>
      <div className="lg:mx-auto lg:w-1/3">
        <TypoWrapper line="bold">
          <h1 className="my-5">{eventName}</h1>
        </TypoWrapper>
        <TypoWrapper size="small" line="shin">
          <p className="text-left lg:mx-10 mx-2 my-2">{detail}</p>
        </TypoWrapper>
        <div className="flex justify-end gap-5 mt-10 mb-2 mx-3">
          <LinkTo href={`/event/${id}/admin/edit`}>編集する</LinkTo>
          <LinkTo href={`/event/${id}/admin/close`}>締切にする</LinkTo>
        </div>
      </div>
    </>
  );
};
