#  by zhouwenzhe at 2023/5/17
from threading import Lock, Thread


def printFirst():
    print("first")


def printSecond():
    print("second")


def printThird():
    print("third")


class Foo:
    def __init__(self):
        self.firstJobDone = Lock()
        self.secondJobDone = Lock()
        self.firstJobDone.acquire()
        self.secondJobDone.acquire()
        pass

    def first(self, printFirst: 'Callable[[], None]') -> None:
        # printFirst() outputs "first". Do not change or remove this line.
        printFirst()
        self.firstJobDone.release()

    def second(self, printSecond: 'Callable[[], None]') -> None:
        # printSecond() outputs "second". Do not change or remove this line.
        with self.firstJobDone:
            printSecond()
            self.secondJobDone.release()

    def third(self, printThird: 'Callable[[], None]') -> None:
        # printThird() outputs "third". Do not change or remove this line.
        with self.secondJobDone:
            printThird()


if __name__ == '__main__':
    foo = Foo()

    thread1 = Thread(target=foo.first, args=(printFirst,))
    thread2 = Thread(target=foo.second, args=(printSecond,))
    thread3 = Thread(target=foo.third, args=(printThird,))

    thread1.start()
    thread3.start()
    thread2.start()

    thread1.join()
    thread3.join()
    thread2.join()
