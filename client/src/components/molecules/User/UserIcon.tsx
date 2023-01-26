import Image from 'next/image';

type TProps = {
  image: string;
  isParticipate: boolean | undefined;
};
export const UserIcon = ({ isParticipate, image }: TProps) => {
  return (
    <>
      {image ? (
        <Image
          src={image}
          alt="user-icon"
          width={50}
          height={50}
          className="rounded-full mx-auto"
        />
      ) : (
        <div
          className={`${
            isParticipate ? 'w-8 h-8' : 'w-12 h-12'
          } rounded-full bg-orange-500 m-auto my-1`}
        ></div>
      )}
    </>
  );
};
