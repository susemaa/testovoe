#include <iostream>
#include <fstream>
#include <string>
#include <cmath>
using namespace std;

int amount_lines = -1, amount_columns = 1;
int **is_calculating;

string get_value(string** whole, string a, int x),
calculate(string a, string b, char operation),
unformula(string** whole, int line, int column),
calculate(string a, string b, char operation);

bool is_formula(string a),
is_number(string a);

int to_int(string a);

//returns true (1) string a is a formula
bool is_formula(string a)
{
	if (a[0] == '=')
		return 1;
	else
		return 0;
}

//transforms string a to int 
int to_int(string a)
{
	int res = 0, j = 0;
	if (a[0] == '-')
		for (int i = a.size() - 1; i > 0; i--)
		{
			res -= (a[i] - 48) * pow(10, j);
			j++;
		}
	else
		for (int i = a.size() - 1; i >= 0; i--)
		{
			res += (a[i] - 48) * pow(10, j);
			j++;
		}
	return res;
}

//returns string value res which is resulf of a *operation* b (e.g. a=12,b=14,op=- -> res=-2)
string calculate(string a, string b, char operation)
{
	int value[3];
	string res;
	value[0] = to_int(a);
	value[1] = to_int(b);
	if (operation == '+')
		value[2] = value[0] + value[1];
	else if (operation == '-')
		value[2] = value[0] - value[1];
	else if (operation == '*')
		value[2] = value[0] * value[1];
	else if (operation == '/')
		if (value[1] == 0)
		{
			cout << "we ran into division by 0 :(" << endl;
			return "";
		}
		else
			value[2] = value[0] / value[1];
	res = to_string(value[2]);
	return res;
}

//returns true (1) if name is one of column names of input csv file
bool is_cellname(string** whole, string name)
{
	string tmp;
	for (int i = 0; i < name.size(); i++)
	{
		tmp += name[i];
		for (int j = 0; j < amount_columns; j++)
			if (tmp == whole[0][j])
				return 1;
	}
	return 0;

}

//if we dont have the cellname in x/y we have to check whether it is a number; terminate if not
bool is_number(string a)
{
	char tmp;
	for (int i = 0; i < a.size(); i++)
	{
		tmp = a[i];
		if (tmp == '0' || tmp == '1' || tmp == '2' || tmp == '3' || tmp == '4' || tmp == '5' || tmp == '6' || tmp == '7' || tmp == '8' || tmp == '9')
			continue;
		else return 0;
	}
	return 1;
}

//recursive function which solves series of formulas until the first one wouldn`t be solved (e.g. a12=b12+2; b12=c12+2; c12=d12+2; d12=2+2 -> d12=4 -> c12=6 -> b12=8 -> a12=10)
string unformula(string** whole, int line, int column)
{
//	cout<<endl;
//	for(int i=0;i<amount_lines;i++)
//	{	
//		for(int j=0;j<amount_columns;j++)
//			cout<<i<<j<<" "<<is_calculating[i][j]<<" "; 
//		cout<<endl;
//	}
	////////////////////////////////
	if(is_calculating[line][column] == 1)
	{
		cout<<"we got into eternal loop :("<<endl;
		return "";
	}
	is_calculating[line][column] = 1;
	////////////////////////////////
	char ch, sign;
	int i = 1, z;
	int new_line = 0, new_column = 0, jj = 0;
	string x, y, name, new_line_str;
	for (i; ; i++)
	{
		ch = whole[line][column][i];
		if (ch == '+' || ch == '-' || ch == '*' || ch == '/')
		{
			sign = ch;
			i++;
			break;
		}
		x += whole[line][column][i];
	}
	for (i; i < whole[line][column].size(); i++)
	{
		y += whole[line][column][i];
	}
	
	if (is_cellname(whole, x) == 0)
		if (is_number(x) == 0)
		{
			cout << "we got wrong cell name" << endl;
			return "";
		}
	if (is_cellname(whole, y) == 0)
		if (is_number(y) == 0)
		{
			cout << "we got wrong cell name" << endl;
			return "";
		}
	
	//for x
	while (is_cellname(whole, x))
	{
		z = 0;
		new_column = -1;

		for (int j = 0; j < x.size(); j++) //getting name and number of cell
		{
			if (new_column == -1)
			{
				name += x[j];
				for (int ij = 0; ij < amount_columns; ij++)
				{
					if (name == whole[0][ij] != 0) // if equal
						new_column = ij;
				}
			}
			else
				new_line_str += x[j];
		}
		new_line = -1;
		for (int j = 0; j < amount_lines; j++)
		{
			if (new_line_str == whole[j][0]) //if equal
				new_line = j;
		}
		if (new_line == -1)
		{
			cout << "we got out of line number :(" << endl;
			return "";
		}

		if ((is_formula(whole[new_line][new_column]))) //is a formula
		{
			whole[new_line][new_column] = unformula(whole, new_line, new_column);
		}
		else
		{
			x = whole[new_line][new_column];
		}
		name.clear();
		new_line_str.clear();
	}

	//for y
	while (is_cellname(whole, y))
	{
		z = 0;
		new_column = -1;

		for (int j = 0; j < y.size(); j++) //getting name and number of cell
		{
			if (new_column == -1)
			{
				name += y[j];
				for (int ij = 0; ij < amount_columns; ij++)
				{
					if (name == whole[0][ij] != 0) // if equal
						new_column = ij;
				}
			}
			else
				new_line_str += y[j];
		}
		new_line = -1;
		for (int j = 0; j < amount_lines; j++)
		{
			if (new_line_str == whole[j][0]) //if equal
				new_line = j;
		}
		if (new_line == -1)
		{
			cout << "we got out of line number :(" << endl;
			return "";
		}

		if ((is_formula(whole[new_line][new_column]))) //is a formula
		{
			whole[new_line][new_column] = unformula(whole, new_line, new_column);
		}
		else
		{
			y = whole[new_line][new_column];
		}
		name.clear();
		new_line_str.clear();
	}
	
	
	string res = calculate(x, y, sign);
	
	return (res);
	
	terminate:
	return "";
}

int main()
{
	//reading filename, open file and check if it opened correctly
	ifstream file;
	string csv_name;
	while (!file.is_open())
	{
		cout << "enter correct csv file name (dont forget about .csv)" << endl;
		cin >> csv_name;
		file.open(csv_name);
		if (file.is_open())
		{
			cout << "file is opened correctly" << endl;
		}
	}

	//counting amount of lines and columns 
	char tmp_ch;
	while (!file.eof()) {
		file.get(tmp_ch);
		if (amount_lines == -1)
			if (tmp_ch == ',')
				amount_columns++;
		if (tmp_ch == '\n')
			amount_lines++;
	}
	cout << "amount of lines: " << amount_lines << endl;
	cout << "amount of columns: " << amount_columns << endl;
	
	
	//allocate memory for array (whole_file) which has every value of csv file
	string** whole_file;
	whole_file = new string * [amount_lines];
	for (int i = 0; i < amount_lines; i++)
		whole_file[i] = new string[amount_columns];
	
	
	//is_calculating is array that consists condition of formula, whether it is calculating or not 
	is_calculating  = new int *[amount_lines];
	for(int i = 0; i < amount_lines; i++)
		is_calculating[i] = new int [amount_columns];
	for (int ii = 0; ii < amount_lines; ii ++)
		for(int jj = 0; jj < amount_columns; jj++)
			is_calculating[ii][jj] = 0;
	
	
	//filling whole_file array
	string tmp_str;
	int cn = 0, ln = 0; //column_number, line_number
	file.clear();
	file.seekg(0);
	while (!file.eof())
	{
		file.get(tmp_ch);
		if (tmp_ch == ',' || tmp_ch == '\n')
		{
			whole_file[ln][cn] += tmp_str;
			cn++;
			tmp_str.clear();
			if (tmp_ch == '\n')
			{
				cn = 0;
				ln++;
				if (ln == amount_lines)
					break;
			}
		}
		else
			tmp_str += tmp_ch;
	}
	
	
	//output source table
	cout << endl << "whole file before calculatings:" << endl << endl;
	for (int i = 0; i < amount_lines; i++)
	{
		for (int j = 0; j < amount_columns; j++)
		{
			cout << whole_file[i][j];
			if (j < amount_columns - 1)
				cout << ",";
		}
		cout << endl;
	}
	
	
	//output final table
	cout << endl << "whole file after calculatings:" << endl << endl;
	for (int i = 0; i < amount_lines; i++)
	{
		for (int j = 0; j < amount_columns; j++)
		{
			if (is_formula(whole_file[i][j])) //is a formula
			{
				whole_file[i][j] = unformula(whole_file, i, j);
				if (whole_file[i][j] == "")
					goto finish;
				for (int ii = 0; ii < amount_lines; ii ++)
					for(int jj = 0; jj < amount_columns; jj++)
						is_calculating[ii][jj] = 0;
			}
			cout << whole_file[i][j];
			if (j < amount_columns - 1)
				cout << ",";
		}
		cout << endl;
	}
	
	//delete all allocated memory and close file
	finish:
	
	for (int i = 0; i < amount_lines; i++)
	{
		delete[]whole_file[i];
	}
	delete[]whole_file;
	
	for (int i = 0; i < amount_lines; i++)
	{
		delete[]is_calculating[i];
	}
	delete[]is_calculating;
	
	file.close();
	system("pause");
	return 0;
}