#include <stdio.h>
#include <stdbool.h>
bool check(int resources, int need[resources], int available[resources]);
void getsafesequence(int processes, int resoures, int allocated[processes][resoures], int max[processes][resoures], int need[processes][resoures], int *available);
void check_request(int process_number, int processes, int resources, int request[resources], int allocated[processes][resources], int max[processes][resources], int need[processes][resources], int available[resources]);
int main()
{
    int processes, resources;
    printf("Enter number of processes : ");
    scanf("%d", &processes);
    printf("Enter number of resoures : ");
    scanf("%d", &resources);
    int allocated[processes][resources], max[processes][resources];
    printf("Enter Allocation Matrix\n");
    for (int i = 0; i < processes; i++)
    {
        for (int j = 0; j < resources; j++)
        {
            scanf("%d", &allocated[i][j]);
        }
    }
    printf("Enter Max Matrix\n");
    for (int i = 0; i < processes; i++)
    {
        for (int j = 0; j < resources; j++)
        {
            scanf("%d", &max[i][j]);
        }
    }
    int avaliable[resources];
    for (int i = 0; i < resources; i++)
    {
        printf("\nenter available of resoures %c: ", (i + 65));
        scanf("%d", &available[i]);
    }
    printf("\nThe Number Of Instances Of Resource Present In The System Under Each Type Of Resource are :\n");
    int instances[resources];
    for (int i = 0; i < resources; i++)
    {
        instances[i] = 0;
    }
    for (int i = 0; i < processes; i++)
    {
        for (int j = 0; j < resources; j++)
        {
            instances[j] += allocated[i][j];
        }
    }
    for (int i = 0; i < resources; i++)
    {
        instances[i] += available[i];
    }
    for (int i = 0; i < resources; i++)
    {
        printf("%c = %d\n", (i + 65), instances[i]);
    }
    printf("\n need matrix is \n");
    int need[processes][resources];
    for (int i = 0; i < processes; i++)
    {
        for (int j = 0; j < resources; j++)
        {
            need[i][j] = max[i][j] - allocated[i][j];
            printf("%d ", need[i][j]);
        }
        printf("\n");
    }
    int current_available[processes];
    for (int i = 0; i < processes; i++)
    {
        current_available[i] = available[i];
    }
    getSafeSequence(processes, resources, allocated, max, need, current_available);
    printf("\n\nIf a request from process p1 arrives for (1,1,0,0), can the request be granted?");
    int request1[4];
    request1[0] = 1;
    request1[1] = 1;
    int current_available1[processes];
    for (int i = 0; i < processes; i++)
    {
        current_available1[i] = available[i];
    }
    check_request(1, processes, resources, request1, allocated, max, need, current_available1);
    printf("\n\nIf a request from process p4 arrives for (0,0,2,0), can the request be granted?\n");
    int request2[4];
    request2[2] = 2;
    int current_available2[processes];
    for (int i = 0; i < processes; i++)
    {
        current_available2[i] = available[i];
    }
    check_request(4, processes, resources, request2, allocated, max, need, current_available2);
    void getSafeSequence(int processes, int resources, int allocated[processes][resources], int max[processes][resources], int need[processes][resources], int available[resources])
    {
        int computed = 0;
        int computed_order[processes];
        int pointer_to_computed = 0;
        bool processed[processes];
        for (int i = 0; i < processes; i++)
        {
            processed[i] = false;
        }
        while (computed < processes)
        {
            bool any_process_computed = false;
            for (int i = 0; i < processes; i++)
            {
                if (processed[i])
                {
                    continue;
                }
                if (check(resources, need[i], available))
                {
                    for (int j = 0; j < resources; j++)
                    {
                        available[j] += allocated[i][j];
                    }
                }
                processed[i] = true;
                any_process_computed = true;
                computed_order[pointer_to_computed++] = i;
                computed += 1;
            }
        }
        if (!any_process_computed)
        {
            break;
        }
    }
    if (computed == processes)
    {
        printf("\nThe System is in safe state and the safe sequence is :\n");
        for (int i = 0; i < processes; i++)
        {
            printf("P%d ", computed_order[i]);
        }
    }
    else
    {
        printf("The System is not in safe state and the processes will be in deadlock\n");
    }
}
bool check(int resources, int need[resources], int available[resources])
{
    for (int i = 0; i < resources; i++)
    {
        if (need[i] > available[i])
        {
            return false;
        }
    }
    return true;
}
void check_request(int process_number, int processes, int resources, int request[resources], int allocated[processes][resources], int max[processes][resources], int need[processes][resources], int available[resources])
{
    if (check(resources, request, available))
    {
        if (check(resources, request, need[process_number]))
        {

            for (int i = 0; i < resources; i++)
            {
                available[i] -= request[i];
                allocated[process_number][i] += request[i];
                need[process_number][i] = max[process_number][i] - allocated[process_number][i];
            }
            getSafeSequence(processes, resources, allocated, max, need, available);
            return;
        }
    }
    printf("Request cannot be granted\n");
}
