---
id: FlDl9gx2mBtFM41D123aX
title: C2-IntelligentAgents
desc: ''
updated: 1639932292448
created: 1639841130136
---

- [Autonomy](#autonomy)
- [Agents and](#agents-and)
  - [Objects](#objects)
  - [Expert systems](#expert-systems)
  - [Intelligent agents and AI](#intelligent-agents-and-ai)
- [Environment Properties](#environment-properties)
- [Intelligent agents](#intelligent-agents)
  - [Reactive (environment aware)](#reactive-environment-aware)
  - [Pro-active](#pro-active)
  - [Social ability](#social-ability)
  - [Other properties](#other-properties)
- [Intentional System](#intentional-system)
- [Abstract architecture](#abstract-architecture)
  - [Environments](#environments)
  - [Agents](#agents)
  - [System](#system)
- [Intelligent agent properties](#intelligent-agent-properties)
  - [Agents with States](#agents-with-states)
- [Task for Agents](#task-for-agents)
  - [Utility functions](#utility-functions)
  - [Optimal agents](#optimal-agents)
  - [Task environment](#task-environment)

# Autonomy
- Tính tự trị
  - Decisions making
- Adjustable: Decisions handed to a higher authority when this is benefical

# Agents and

## Objects
- Agent = object + **attitude**
  
- Object:
  - States
  - Message passing
  - Methods (operations that may be performed on this state)

- Attitude
  - Autonomous: Decide for themselves whether or not to follow other's requests
  - Smart: Flexible (reactive, pro-active, social) behaviour
  - Active: each agent has at least 1 thread of active control

## Expert systems
- Def: "expertise" about some (abstract) domain of discourse (bàn luận)
- Agent is **aware** of the world

## Intelligent agents and AI
- Agent can be built to operate in **a limited domain**
- A useful agent is not needed to solve all the AI's problems

# Environment Properties
- Observable: fully or partially
  - Fully: The agent can obtain complete, accurate, up-to-date information about the environment's state
- Deterministic or non-deterministic
  - Deterministic: Any action has a single **guaranteed effect** (no uncertainty)
- Static or dynamic
  - Static: Remain **unchanged** except by the agent's action
- Discrete or continuous
  - Discrete: **a fixed, finite number of actions** and percepts in it
- Episodic or non-episodic
  - Phân đoạn
  - Episodic environment
    - The performance of agent is **dependent on a number of discrete eptisodes**
    - No linkages between different scenarios
- Real time
  - **Time** plays a part in **evaluating** an agents performance

# Intelligent agents

## Reactive (environment aware)
- Phản ứng nhanh
- A reactive system is one that
  - Maintains an ongoing interaction with its environment
  - **Responds to changes** that occurs in it

## Pro-active
- Chuyên nghiệp
- Generating and attempting to achieve goals

## Social ability
- Taking **others** into account
- Cooperation: working **together** to achieve a **shared goal**
- Coordination: Managing the **interdependencies** between activities (sharing resources)
- Negotiation: To reach **agreements** on matters of **common interest**
  - Offer
  - Counter offer

## Other properties
- Mobility: moving
- Rationality (hợp lý): Act to achieve goals
- Veracity (xác thực): know the communication failures
- Benevolence (nhân từ): to help or not to help
- Learning/Adaption

# Intentional System
- Folk psychology:
  - **Human behaviour is predicted** and explainend through the attribution of **attitudes**.
  - Attitudes = intentional notions

# Abstract architecture
- The world 
  - Finite set E of discrete, instantaneous (tức thời) **states** 
    - ![](/assets/images/2021-12-19-17-09-25.png)
- Agents have a set of **possible actions** to transform the state of the world
  - ![](/assets/images/2021-12-19-17-10-59.png)
- A **run** of an agent is a **sequence** of interleaved (xen kẽ) _world states and actions_
  - ![](/assets/images/2021-12-19-17-12-45.png)  
  - Runs can end with a state or an action
    - ![](/assets/images/2021-12-19-17-21-18.png)

## Environments
- Properties
  - History dependent
  - Non-determistic
- State transformer function: environment's behaviour
  - ![](/assets/images/2021-12-19-17-24-49.png)
- ![](/assets/images/2021-12-19-17-45-14.png)
  - E: set of states
  - e[0]: initial state
  - T: state transformer function

## Agents
- ![](/assets/images/2021-12-19-17-47-19.png)
  - Agent = function which maps **runs to actions**
  - Ag: the set of all agents

## System
- ![](/assets/images/2021-12-19-17-49-38.png)
- A system = an agent + an environment
- Associate with a set of **possible runs**
- ![](/assets/images/2021-12-19-17-55-43.png)

# Intelligent agent properties
- Deliberative (chủ ý): Agent will reach a **different decision** when it reach the **same state** by **different routes** 
- Purely reactive
  - Without history references
  - ![](/assets/images/2021-12-19-18-00-52.png) (state to actions)
  - Reactive agent
    - Always do the **same thing** in the **same state**

## Agents with States
- ![](/assets/images/2021-12-19-22-43-54.png)
- Agent's internal **data structure**
  - Record information about the environment state & history
- **See**: observe the environment
  - ![](/assets/images/2021-12-19-22-44-54.png)
  - Output is a percept (nhận thức)
- **Action**: decision making
  - ![](/assets/images/2021-12-19-22-48-24.png)
  - Internal states -> actions
- **Next**: 
  - ![](/assets/images/2021-12-19-22-51-50.png)
  - Internal States + percept => new internal states
  - **Updates** the agent's view when it gets a **new percept**

# Task for Agents

## Utility functions
- **Rewarding** agent with the state it brings out
- A task specification
  - A real number (reward) with every environment state
  - ![](/assets/images/2021-12-19-22-56-27.png)
- Local utility functions
  - Assigning utilities to **individual states**
  - Difficult to specify a **long term view**
  - **Reinforcement** learning: A **discount** for states later on
- Sequential decision making
  - Utility gained depends on the route taken
- Utilities over Runs
  - ![](/assets/images/2021-12-19-23-01-28.png)
  - Assign utility for runs
- Expected utility
  - Probality that run _r_ occurs when agent _Ag_ is placed in the enviroment _Env_
    - ![](/assets/images/2021-12-19-23-04-15.png)
  - ![](/assets/images/2021-12-19-23-04-50.png)
    - Utility of each run
    - Possibility of each run

## Optimal agents
- Maximizes expected utility
  - ![](/assets/images/2021-12-19-23-08-11.png)
- Do the best **on average**
- Bounded optimal agents
  - Some agents **can not be implemented**
  - Include only those agents tthat can be implemented on **machine m**
    - ![](/assets/images/2021-12-19-23-11-13.png)
  - ![](/assets/images/2021-12-19-23-10-54.png)

## Task environment
- Predicate task specification: Assign 0/1 (true/false | succeed/failed) to a run
  - ![](/assets/images/2021-12-19-23-13-26.png)
- Task environment:
  - A pair ![](/assets/images/2021-12-19-23-14-41.png)
    - ![](/assets/images/2021-12-19-23-14-17.png)
    - Environment
    - Task specification
  - Specifies
    - The properties of [the system](#system) the agent will inhabit (occupy)
    - The **criteria** by which an agent will be **judged**
- A set of runs of agent _Ag_ in environment _Env_ that satisfy _$\psi$_
  - ![](/assets/images/2021-12-19-23-23-52.png) 
  - Agent _Ag_ succeeds in task environment $<Env, \psi>$
    - ![](/assets/images/2021-12-19-23-27-12.png)
    - An agent succeeds if **every run satisfies the specification of the agent**
      - ![](/assets/images/2021-12-19-23-29-32.png)
    - More optimistic: succeed = at least a successful run
      - ![](/assets/images/2021-12-19-23-29-18.png)
- The probability of success
  - Non-deterministic: the state transform function returns **a set of possible states**
  - ![](/assets/images/2021-12-19-23-34-14.png)
    - The sum of probabilities of every run that satisfy $\psi$ when agent _Ag_ is placed in environment _Env_
- Types
  - Achievement: if the agent can **force the environment** into one of the goal states
  - Maintenance: if the agent **never forced into** one of the fail states