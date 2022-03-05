#include <stdio.h>

int main(void)
{
	int a;
	for (int i = 0; i < 100; i++)
	{
		a += i;
	}
	printf("%d\n", a);
	return 0;
}