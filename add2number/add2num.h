#ifndef ADD2NUM_H
#define ADD2NUM_H

#include <string>
#include <vector>
using namespace std;


 struct node
{
    int number;
    node *next;
    bool head_flag=false;
    bool up_number = false;
};

void add2number(string number_one, string number_two);

class add2number_class
{
    public:
        add2number_class(){}
        ~add2number_class();
        

        void link_list_create(const string &number_one,const string &number_two);
        void sum_the_two_number();
        void print_the_link_list(node *headeer_node);
        void update_current(node *&current_one, node *&current_two, node *&current);
        void make_sum(node *&current_one, node *&current_two, node *&new_node, bool up);
        node* block_generator(const string &number);
        node *header_generator(int &number);
    private:       
        vector<node*> node_header;
        
};

#endif