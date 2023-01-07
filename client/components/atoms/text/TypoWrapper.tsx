import { ReactNode } from 'react';
import { match } from 'ts-pattern';

type TProps = {
  children: ReactNode;
  line?: 'shin' | 'bold';
  size?: 'small' | 'large';
};
export const TypoWrapper = ({ line, size, children }: TProps) => {
  const decide_line = ({ line }: Pick<TProps, 'line'>) => {
    return match(line)
      .with('shin', () => 'font-thin')
      .with('bold', () => 'font-bold')
      .with(undefined, () => 'font-normal')
      .exhaustive();
  };
  const decide_size = ({ size }: Pick<TProps, 'size'>) => {
    return match(size)
      .with('small', () => 'text-sm')
      .with('large', () => 'text-2xl')
      .with(undefined, () => 'text-base')
      .exhaustive();
  };
  return (
    <div className={`${decide_line({ line })} ${decide_size({ size })}`}>
      {children}
    </div>
  );
};
