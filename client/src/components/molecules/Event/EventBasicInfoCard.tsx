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
      <div className="lg:mx-auto">
        <TypoWrapper line="bold" size="so-large">
          <h1 className="my-5">{eventName}</h1>
        </TypoWrapper>
        <TypoWrapper line="shin">
          <p className="text-left lg:mx-44 mx-5 py-5">{detail}</p>
        </TypoWrapper>
        <div className="flex justify-end md:gap-5 gap-1 mt-10 mb-2 mx-3">
          <LinkTo href={`/event/${id}/`} borderNone>
            詳細
          </LinkTo>
          <LinkTo href={`/event/${id}/admin/edit`} borderNone>
            編集
          </LinkTo>
          <LinkTo href={`/event/${id}/admin/close`} borderNone>
            締切
          </LinkTo>
        </div>
      </div>
    </>
  );
};
