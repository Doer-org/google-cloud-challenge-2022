import { TypoWrapper } from './TypoWrapper';

type TProps = {
  text: string;
  full?: boolean | undefined;
};
const Arrow = () => {
  return (
    <span
      className="before:absolute before:rotate-[180deg] before:-bottom-1 before:left-[50%] before:-translate-x-[50%] before:content-['']
before:border-t-[10%] before:border-b-white before:border-[10px] before:border-transparent"
    ></span>
  );
};
export const Comment = ({ text, full }: TProps) => {
  return (
    <>
      {full ? (
        <div className="relative py-3 w-screen mx-auto">
          <div className="bg-white text-black rounded-md comment p-5 w-2/3 md:w-1/3 max-h-40 overflow-y-scroll mx-auto text-left">
            <TypoWrapper line="shin">
              <p className="break-all">{text}</p>
            </TypoWrapper>
          </div>
          <Arrow />
        </div>
      ) : (
        <div className="relative py-3">
          <div className="bg-white text-black rounded-md comment p-1 hover:scale-110 transition cursor-pointer">
            <TypoWrapper size="small" line="shin">
              <p className="overflow-x-scroll w-20 whitespace-nowrap">
                {text}...
              </p>
            </TypoWrapper>
          </div>
          <Arrow />
        </div>
      )}
    </>
  );
};
