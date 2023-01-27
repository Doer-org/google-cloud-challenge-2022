import Image from 'next/image';
import { TypoWrapper } from '../../atoms/text/TypoWrapper';
import style from '../../../styles/title.module.css';
export const TopPage = () => {
  return (
    <div>
      <div className="grid grid-cols-1 place-items-center">
        <Image
          src={'/logo.png'}
          alt={'top'}
          width={'230'}
          height={'300'}
          className=""
        />
        <TypoWrapper line="bold">
          <h1 className={style.title}>
            すきま時間を<span className=" text-yellow-500">”すきーま”</span>で
          </h1>
          <h1 className={style.title}>有効活用しませんか？</h1>
        </TypoWrapper>
        <Image
          src={'/girl.png'}
          alt={'top'}
          width={'150'}
          height={'300'}
          className=""
        />
      </div>
    </div>
  );
};
