import { Comment } from '../../atoms/text/Comment';

type TProps = {
  comment: string | undefined;
  full?: boolean;
};
export const UserComment = ({ comment, full }: TProps) => {
  return (
    <>
      {comment ? (
        <>
          {full ? (
            <Comment text={comment} full />
          ) : (
            <Comment text={comment.slice(0, 3)} />
          )}
        </>
      ) : (
        <> {full ? <Comment text="..." full /> : <Comment text="..." />}</>
      )}
    </>
  );
};
