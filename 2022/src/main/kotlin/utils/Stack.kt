package utils

class Stack<T> {
    private var stack = mutableListOf<T>()

    fun size(): Int = stack.size;
    fun top(): T = stack.last()

    fun insert(s: T) = stack.add(0, s)
    fun put(s: T) = stack.add(s)

    fun putInOrder(elements: MutableList<T>) {
        val success = stack.addAll(elements)
    }

    fun pop(): T = stack.removeLast()

    fun list() = stack

    fun popInOrder(amount: Int): MutableList<T> {
        val poppedItems = stack.subList(stack.size - amount, stack.size)
        stack = stack.subList(0, stack.size - amount)

        return poppedItems
    }
}