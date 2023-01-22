type TProps = {
  isParticipate: boolean | undefined;
};
export const UserIcon = ({ isParticipate }: TProps) => {
  return (
    <div
      className={`${
        isParticipate ? 'w-8 h-8' : 'w-12 h-12'
      } rounded-full bg-orange-500 m-auto my-1`}
    ></div>
  );
};
