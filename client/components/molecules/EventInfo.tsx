import { TypoWrapper } from '../atoms/TypoWrapper';
import { Map } from '../atoms/Map';
export const EventInfo = () => {
  return (
    <div className="m-auto my-5">
      <div className="my-5 m-auto">
        <div className="w-14 h-14 rounded-full bg-orange-500 m-auto my-1"></div>
        <TypoWrapper size="small" line="shin">
          <p>userName</p>
        </TypoWrapper>
      </div>
      <div>
        <TypoWrapper size="large" line="bold">
          <h1>今からちょっとご飯</h1>
        </TypoWrapper>
        <TypoWrapper size="small" line="shin">
          <p className="text-left mx-10 my-2">
            同志社周りで今からラーメン行ける人いませんか？？あくたがわ行こうと思ってます！！
          </p>
        </TypoWrapper>
      </div>
      <div className="m-10">
        <Map />
      </div>
    </div>
  );
};
