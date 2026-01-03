This is a very simple python evaluator for an interpreter, It give me a better understanding of how evaluation happens in an interpreter

> NOTE: This assumes the parser forms an AST and not bytecode, We are recursively travasing the tree and calling the eval function

```python
class Node:
    def __init__(self, value=None):
        self.value = value
        self.left = None
        self.right = None

def evaluate(node):
    if isinstance(node, Node):
        if node.value is not None:
            return node.value  # Base case: a constant or variable reference
        else:
            left_value = evaluate(node.left)
            right_value = evaluate(node.right)

            if node.op == '+':
                return left_value + right_value
            elif node.op == '-':
                return left_value - right_value
            elif node.op == '*':
                return left_value * right_value
            elif node.op == '/':
                return left_value / right_value
            else:
                raise ValueError(f"Unsupported operation: {node.op}")
    else:
        raise TypeError("Invalid type for the node")
# Example AST
root = Node(5)  # Constant value
add_node = Node('+')
add_node.left = root
add_node.right = Node(3)

# Evaluate the AST
result = evaluate(add_node)
print(result)  # Output: 8
```

But for bat, we are modeling the eval returns into an Object struct that would be passed around
