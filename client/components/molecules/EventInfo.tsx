import { TypoWrapper } from '../atoms/TypoWrapper';
import { Map } from '../atoms/Map';
export const EventInfo = () => {
  return (
    <>
      <div className="flex justify-center items-center flex-col">
        <div className="w-8 h-8 rounded-full bg-slate-400 m-5 z-20"></div>
        <div className="flex justify-center">
          <div className="z-10">
            <div className="border border-black w-24 -rotate-45"></div>
            <div className="w-3 h-3 rounded-full bg-slate-400 mr-auto ml-2 mt-6 z-20 relative"></div>
          </div>
          <div className="z-10">
            <div className="border border-black w-24 rotate-45"></div>
            <div className="w-3 h-3 rounded-full bg-slate-400 ml-auto mr-2 mt-6 z-20 relative"></div>
          </div>
        </div>
      </div>

      <div className="m-5 border-accent_border border-8 shadow-2xl bg-origin_depth rounded-md -mt-3">
        <div className="lg:m-3 m-1 lg:p-3 p-1 border-white border rounded-md">
          <div className="my-5 m-auto">
            <div className="w-14 h-14 rounded-full bg-orange-500 m-auto my-1"></div>
            <TypoWrapper size="small" line="shin">
              <p>userName</p>
            </TypoWrapper>
          </div>
          <div className="lg:mx-auto lg:w-1/3">
            <TypoWrapper size="large" line="bold">
              <h1>今からちょっとご飯</h1>
            </TypoWrapper>
            <TypoWrapper size="small" line="shin">
              <p className="text-left lg:mx-10 mx-2 my-2">
                同志社周りで今からラーメン行ける人いませんか？？あくたがわ行こうと思ってます！！
              </p>
            </TypoWrapper>
          </div>
          <div className="lg:m-10 m-3">
            <Map />
          </div>
        </div>
      </div>
    </>
  );
};
