// Dafny program main.dfy compiled into JavaScript
// Copyright by the contributors to the Dafny Project
// SPDX-License-Identifier: MIT

const BigNumber = require('bignumber.js');
BigNumber.config({ MODULO_MODE: BigNumber.EUCLID })
let _dafny = (function() {
  let $module = {};
  $module.areEqual = function(a, b) {
    if (typeof a === 'string' && b instanceof _dafny.Seq) {
      // Seq.equals(string) works as expected,
      // and the catch-all else block handles that direction.
      // But the opposite direction doesn't work; handle it here.
      return b.equals(a);
    } else if (typeof a === 'number' && BigNumber.isBigNumber(b)) {
      // This conditional would be correct even without the `typeof a` part,
      // but in most cases it's probably faster to short-circuit on a `typeof`
      // than to call `isBigNumber`. (But it remains to properly test this.)
      return b.isEqualTo(a);
    } else if (typeof a !== 'object' || a === null || b === null) {
      return a === b;
    } else if (BigNumber.isBigNumber(a)) {
      return a.isEqualTo(b);
    } else if (a._tname !== undefined || (Array.isArray(a) && a.constructor.name == "Array")) {
      return a === b;  // pointer equality
    } else {
      return a.equals(b);  // value-type equality
    }
  }
  $module.toString = function(a) {
    if (a === null) {
      return "null";
    } else if (typeof a === "number") {
      return a.toFixed();
    } else if (BigNumber.isBigNumber(a)) {
      return a.toFixed();
    } else if (a._tname !== undefined) {
      return a._tname;
    } else {
      return a.toString();
    }
  }
  $module.escapeCharacter = function(cp) {
    let s = String.fromCodePoint(cp.value)
    switch (s) {
      case '\n': return "\\n";
      case '\r': return "\\r";
      case '\t': return "\\t";
      case '\0': return "\\0";
      case '\'': return "\\'";
      case '\"': return "\\\"";
      case '\\': return "\\\\";
      default: return s;
    };
  }
  $module.NewObject = function() {
    return { _tname: "object" };
  }
  $module.InstanceOfTrait = function(obj, trait) {
    return obj._parentTraits !== undefined && obj._parentTraits().includes(trait);
  }
  $module.Rtd_bool = class {
    static get Default() { return false; }
  }
  $module.Rtd_char = class {
    static get Default() { return 'D'; }  // See CharType.DefaultValue in Dafny source code
  }
  $module.Rtd_codepoint = class {
    static get Default() { return new _dafny.CodePoint('D'.codePointAt(0)); }
  }
  $module.Rtd_int = class {
    static get Default() { return BigNumber(0); }
  }
  $module.Rtd_number = class {
    static get Default() { return 0; }
  }
  $module.Rtd_ref = class {
    static get Default() { return null; }
  }
  $module.Rtd_array = class {
    static get Default() { return []; }
  }
  $module.ZERO = new BigNumber(0);
  $module.ONE = new BigNumber(1);
  $module.NUMBER_LIMIT = new BigNumber(0x20).multipliedBy(0x1000000000000);  // 2^53
  $module.Tuple = class Tuple extends Array {
    constructor(...elems) {
      super(...elems);
    }
    toString() {
      return "(" + arrayElementsToString(this) + ")";
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static Default(...values) {
      return Tuple.of(...values);
    }
    static Rtd(...rtdArgs) {
      return {
        Default: Tuple.from(rtdArgs, rtd => rtd.Default)
      };
    }
  }
  $module.Set = class Set extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Set.Empty;
    }
    toString() {
      return "{" + arrayElementsToString(this) + "}";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Set();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new Set();
      for (let k of elmts) {
        s.add(k);
      }
      return s;
    }
    contains(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i], k)) {
          return true;
        }
      }
      return false;
    }
    add(k) {  // mutates the Set; use only during construction
      if (!this.contains(k)) {
        this.push(k);
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        if (!other.contains(e)) {
          return false;
        }
      }
      return true;
    }
    get Elements() {
      return this;
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = Set.of(...this);
        for (let k of that) {
          s.add(k);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new Set();
        for (let k of this) {
          if (that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length == 0 || that.length == 0) {
        return this;
      } else {
        let s = new Set();
        for (let k of this) {
          if (!that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      for (let k of this) {
        if (that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsSubsetOf(that) {
      if (that.length < this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      if (that.length <= this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    get AllSubsets() {
      return this.AllSubsets_();
    }
    *AllSubsets_() {
      // Start by putting all set elements into a list, but don't include null
      let elmts = Array.of(...this);
      let n = elmts.length;
      let which = new Array(n);
      which.fill(false);
      let a = [];
      while (true) {
        yield Set.of(...a);
        // "add 1" to "which", as if doing a carry chain.  For every digit changed, change the membership of the corresponding element in "a".
        let i = 0;
        for (; i < n && which[i]; i++) {
          which[i] = false;
          // remove elmts[i] from a
          for (let j = 0; j < a.length; j++) {
            if (_dafny.areEqual(a[j], elmts[i])) {
              // move the last element of a into slot j
              a[j] = a[-1];
              a.pop();
              break;
            }
          }
        }
        if (i === n) {
          // we have cycled through all the subsets
          break;
        }
        which[i] = true;
        a.push(elmts[i]);
      }
    }
  }
  $module.MultiSet = class MultiSet extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return MultiSet.Empty;
    }
    toString() {
      let s = "multiset{";
      let sep = "";
      for (let e of this) {
        let [k, n] = e;
        let ks = _dafny.toString(k);
        while (!n.isZero()) {
          n = n.minus(1);
          s += sep + ks;
          sep = ", ";
        }
      }
      s += "}";
      return s;
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new MultiSet();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new MultiSet();
      for (let e of elmts) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    static FromArray(arr) {
      let s = new MultiSet();
      for (let e of arr) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    cardinality() {
      let c = _dafny.ZERO;
      for (let e of this) {
        let [k, n] = e;
        c = c.plus(n);
      }
      return c;
    }
    clone() {
      let s = new MultiSet();
      for (let e of this) {
        let [k, n] = e;
        s.push([k, n]);  // make sure to create a new array [k, n] here
      }
      return s;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return _dafny.ZERO;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return !this.get(k).isZero();
    }
    add(k, n) {
      let i = this.findIndex(k);
      if (i === this.length) {
        this.push([k, n]);
      } else {
        let m = this[i][1];
        this[i] = [k, m.plus(n)];
      }
    }
    update(k, n) {
      let i = this.findIndex(k);
      if (i < this.length && this[i][1].isEqualTo(n)) {
        return this;
      } else if (i === this.length && n.isZero()) {
        return this;
      } else if (i === this.length) {
        let m = this.slice();
        m.push([k, n]);
        return m;
      } else {
        let m = this.slice();
        m[i] = [k, n];
        return m;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let e of this) {
        let [k, n] = e;
        let m = other.get(k);
        if (!n.isEqualTo(m)) {
          return false;
        }
      }
      return this.cardinality().isEqualTo(other.cardinality());
    }
    get Elements() {
      return this.Elements_();
    }
    *Elements_() {
      for (let i = 0; i < this.length; i++) {
        let [k, n] = this[i];
        while (!n.isZero()) {
          yield k;
          n = n.minus(1);
        }
      }
    }
    get UniqueElements() {
      return this.UniqueElements_();
    }
    *UniqueElements_() {
      for (let e of this) {
        let [k, n] = e;
        if (!n.isZero()) {
          yield k;
        }
      }
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = this.clone();
        for (let e of that) {
          let [k, n] = e;
          s.add(k, n);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let m = that.get(k);
          if (!m.isZero()) {
            s.push([k, m.isLessThan(n) ? m : n]);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length === 0 || that.length === 0) {
        return this;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let d = n.minus(that.get(k));
          if (d.isGreaterThan(0)) {
            s.push([k, d]);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      let intersection = this.Intersect(that);
      return intersection.cardinality().isZero();
    }
    IsSubsetOf(that) {
      for (let e of this) {
        let [k, n] = e;
        let m = that.get(k);
        if (!n.isLessThanOrEqualTo(m)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      return this.IsSubsetOf(that) && this.cardinality().isLessThan(that.cardinality());
    }
  }
  $module.CodePoint = class CodePoint {
    constructor(value) {
      this.value = value
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      return this.value === other.value
    }
    isLessThan(other) {
      return this.value < other.value
    }
    isLessThanOrEqual(other) {
      return this.value <= other.value
    }
    toString() {
      return "'" + $module.escapeCharacter(this) + "'";
    }
    static isCodePoint(i) {
      return (
        (_dafny.ZERO.isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0xD800))) ||
        (new BigNumber(0xE000).isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0x11_0000))))
    }
  }
  $module.Seq = class Seq extends Array {
    constructor(...elems) {
      super(...elems);
    }
    static get Default() {
      return Seq.of();
    }
    static Create(n, init) {
      return Seq.from({length: n}, (_, i) => init(new BigNumber(i)));
    }
    static UnicodeFromString(s) {
      return new Seq(...([...s].map(c => new _dafny.CodePoint(c.codePointAt(0)))))
    }
    toString() {
      return "[" + arrayElementsToString(this) + "]";
    }
    toVerbatimString(asLiteral) {
      if (asLiteral) {
        return '"' + this.map(c => _dafny.escapeCharacter(c)).join("") + '"';
      } else {
        return this.map(c => String.fromCodePoint(c.value)).join("");
      }
    }
    static update(s, i, v) {
      if (typeof s === "string") {
        let p = s.slice(0, i);
        let q = s.slice(i.toNumber() + 1);
        return p.concat(v, q);
      } else {
        let t = s.slice();
        t[i] = v;
        return t;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static contains(s, k) {
      if (typeof s === "string") {
        return s.includes(k);
      } else {
        for (let x of s) {
          if (_dafny.areEqual(x, k)) {
            return true;
          }
        }
        return false;
      }
    }
    get Elements() {
      return this;
    }
    get UniqueElements() {
      return _dafny.Set.fromElements(...this);
    }
    static Concat(a, b) {
      if (typeof a === "string" || typeof b === "string") {
        // string concatenation, so make sure both operands are strings before concatenating
        if (typeof a !== "string") {
          // a must be a Seq
          a = a.join("");
        }
        if (typeof b !== "string") {
          // b must be a Seq
          b = b.join("");
        }
        return a + b;
      } else {
        // ordinary concatenation
        let r = Seq.of(...a);
        r.push(...b);
        return r;
      }
    }
    static JoinIfPossible(x) {
      try { return x.join(""); } catch(_error) { return x; }
    }
    static IsPrefixOf(a, b) {
      if (b.length < a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
    static IsProperPrefixOf(a, b) {
      if (b.length <= a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
  }
  $module.Map = class Map extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Map.of();
    }
    toString() {
      return "map[" + this.map(maplet => _dafny.toString(maplet[0]) + " := " + _dafny.toString(maplet[1])).join(", ") + "]";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Map();
      }
      return this._empty;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return undefined;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return this.findIndex(k) < this.length;
    }
    update(k, v) {
      let m = this.slice();
      m.updateUnsafe(k, v);
      return m;
    }
    // Similar to update, but make the modification in-place.
    // Meant to be used in the map constructor.
    updateUnsafe(k, v) {
      let m = this;
      let i = m.findIndex(k);
      m[i] = [k, v];
      return m;
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        let [k, v] = e;
        let w = other.get(k);
        if (w === undefined || !_dafny.areEqual(v, w)) {
          return false;
        }
      }
      return true;
    }
    get Keys() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(k);
      }
      return s;
    }
    get Values() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.add(v);
      }
      return s;
    }
    get Items() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(_dafny.Tuple.of(k, v));
      }
      return s;
    }
    Merge(that) {
      let m = that.slice();
      for (let e of this) {
        let [k, v] = e;
        let i = m.findIndex(k);
        if (i == m.length) {
          m[i] = [k, v];
        }
      }
      return m;
    }
    Subtract(keys) {
      if (this.length === 0 || keys.length === 0) {
        return this;
      }
      let m = new Map();
      for (let e of this) {
        let [k, v] = e;
        if (!keys.contains(k)) {
          m[m.length] = e;
        }
      }
      return m;
    }
  }
  $module.newArray = function(initValue, ...dims) {
    return { dims: dims, elmts: buildArray(initValue, ...dims) };
  }
  $module.BigOrdinal = class BigOrdinal {
    static get Default() {
      return _dafny.ZERO;
    }
    static IsLimit(ord) {
      return ord.isZero();
    }
    static IsSucc(ord) {
      return ord.isGreaterThan(0);
    }
    static Offset(ord) {
      return ord;
    }
    static IsNat(ord) {
      return true;  // at run time, every ORDINAL is a natural number
    }
  }
  $module.BigRational = class BigRational {
    static get ZERO() {
      if (this._zero === undefined) {
        this._zero = new BigRational(_dafny.ZERO);
      }
      return this._zero;
    }
    constructor (n, d) {
      // requires d === undefined || 1 <= d
      this.num = n;
      this.den = d === undefined ? _dafny.ONE : d;
      // invariant 1 <= den || (num == 0 && den == 0)
    }
    static get Default() {
      return _dafny.BigRational.ZERO;
    }
    // We need to deal with the special case `num == 0 && den == 0`, because
    // that's what C#'s default struct constructor will produce for BigRational. :(
    // To deal with it, we ignore `den` when `num` is 0.
    toString() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num.toFixed() + ".0";
      }
      let answer = this.dividesAPowerOf10(this.den);
      if (answer !== undefined) {
        let n = this.num.multipliedBy(answer[0]);
        let log10 = answer[1];
        let sign, digits;
        if (this.num.isLessThan(0)) {
          sign = "-"; digits = n.negated().toFixed();
        } else {
          sign = ""; digits = n.toFixed();
        }
        if (log10 < digits.length) {
          let digitCount = digits.length - log10;
          return sign + digits.slice(0, digitCount) + "." + digits.slice(digitCount);
        } else {
          return sign + "0." + "0".repeat(log10 - digits.length) + digits;
        }
      } else {
        return "(" + this.num.toFixed() + ".0 / " + this.den.toFixed() + ".0)";
      }
    }
    isPowerOf10(x) {
      if (x.isZero()) {
        return undefined;
      }
      let log10 = 0;
      while (true) {  // invariant: x != 0 && x * 10^log10 == old(x)
        if (x.isEqualTo(1)) {
          return log10;
        } else if (x.mod(10).isZero()) {
          log10++;
          x = x.dividedToIntegerBy(10);
        } else {
          return undefined;
        }
      }
    }
    dividesAPowerOf10(i) {
      let factor = _dafny.ONE;
      let log10 = 0;
      if (i.isLessThanOrEqualTo(_dafny.ZERO)) {
        return undefined;
      }

      // invariant: 1 <= i && i * 10^log10 == factor * old(i)
      while (i.mod(10).isZero()) {
        i = i.dividedToIntegerBy(10);
       log10++;
      }

      while (i.mod(5).isZero()) {
        i = i.dividedToIntegerBy(5);
        factor = factor.multipliedBy(2);
        log10++;
      }
      while (i.mod(2).isZero()) {
        i = i.dividedToIntegerBy(2);
        factor = factor.multipliedBy(5);
        log10++;
      }

      if (i.isEqualTo(_dafny.ONE)) {
        return [factor, log10];
      } else {
        return undefined;
      }
    }
    toBigNumber() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num;
      } else if (this.num.isGreaterThan(0)) {
        return this.num.dividedToIntegerBy(this.den);
      } else {
        return this.num.minus(this.den).plus(1).dividedToIntegerBy(this.den);
      }
    }
    isInteger() {
      return this.equals(new _dafny.BigRational(this.toBigNumber(), _dafny.ONE));
    }
    // Returns values such that aa/dd == a and bb/dd == b.
    normalize(b) {
      let a = this;
      let aa, bb, dd;
      if (a.num.isZero()) {
        aa = a.num;
        bb = b.num;
        dd = b.den;
      } else if (b.num.isZero()) {
        aa = a.num;
        dd = a.den;
        bb = b.num;
      } else {
        let gcd = BigNumberGcd(a.den, b.den);
        let xx = a.den.dividedToIntegerBy(gcd);
        let yy = b.den.dividedToIntegerBy(gcd);
        // We now have a == a.num / (xx * gcd) and b == b.num / (yy * gcd).
        aa = a.num.multipliedBy(yy);
        bb = b.num.multipliedBy(xx);
        dd = a.den.multipliedBy(yy);
      }
      return [aa, bb, dd];
    }
    compareTo(that) {
      // simple things first
      let asign = this.num.isZero() ? 0 : this.num.isLessThan(0) ? -1 : 1;
      let bsign = that.num.isZero() ? 0 : that.num.isLessThan(0) ? -1 : 1;
      if (asign < 0 && 0 <= bsign) {
        return -1;
      } else if (asign <= 0 && 0 < bsign) {
        return -1;
      } else if (bsign < 0 && 0 <= asign) {
        return 1;
      } else if (bsign <= 0 && 0 < asign) {
        return 1;
      }
      let [aa, bb, dd] = this.normalize(that);
      if (aa.isLessThan(bb)) {
        return -1;
      } else if (aa.isEqualTo(bb)){
        return 0;
      } else {
        return 1;
      }
    }
    equals(that) {
      return this.compareTo(that) === 0;
    }
    isLessThan(that) {
      return this.compareTo(that) < 0;
    }
    isAtMost(that) {
      return this.compareTo(that) <= 0;
    }
    plus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.plus(bb), dd);
    }
    minus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.minus(bb), dd);
    }
    negated() {
      return new BigRational(this.num.negated(), this.den);
    }
    multipliedBy(b) {
      return new BigRational(this.num.multipliedBy(b.num), this.den.multipliedBy(b.den));
    }
    dividedBy(b) {
      let a = this;
      // Compute the reciprocal of b
      let bReciprocal;
      if (b.num.isGreaterThan(0)) {
        bReciprocal = new BigRational(b.den, b.num);
      } else {
        // this is the case b.num < 0
        bReciprocal = new BigRational(b.den.negated(), b.num.negated());
      }
      return a.multipliedBy(bReciprocal);
    }
  }
  $module.EuclideanDivisionNumber = function(a, b) {
    if (0 <= a) {
      if (0 <= b) {
        // +a +b: a/b
        return Math.floor(a / b);
      } else {
        // +a -b: -(a/(-b))
        return -Math.floor(a / -b);
      }
    } else {
      if (0 <= b) {
        // -a +b: -((-a-1)/b) - 1
        return -Math.floor((-a-1) / b) - 1;
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return Math.floor((-a-1) / -b) + 1;
      }
    }
  }
  $module.EuclideanDivision = function(a, b) {
    if (a.isGreaterThanOrEqualTo(0)) {
      if (b.isGreaterThanOrEqualTo(0)) {
        // +a +b: a/b
        return a.dividedToIntegerBy(b);
      } else {
        // +a -b: -(a/(-b))
        return a.dividedToIntegerBy(b.negated()).negated();
      }
    } else {
      if (b.isGreaterThanOrEqualTo(0)) {
        // -a +b: -((-a-1)/b) - 1
        return a.negated().minus(1).dividedToIntegerBy(b).negated().minus(1);
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return a.negated().minus(1).dividedToIntegerBy(b.negated()).plus(1);
      }
    }
  }
  $module.EuclideanModuloNumber = function(a, b) {
    let bp = Math.abs(b);
    if (0 <= a) {
      // +a: a % bp
      return a % bp;
    } else {
      // c = ((-a) % bp)
      // -a: bp - c if c > 0
      // -a: 0 if c == 0
      let c = (-a) % bp;
      return c === 0 ? c : bp - c;
    }
  }
  $module.ShiftLeft = function(b, n) {
    return b.multipliedBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.ShiftRight = function(b, n) {
    return b.dividedToIntegerBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.RotateLeft = function(b, n, w) {  // truncate(b << n) | (b >> (w - n))
    let x = _dafny.ShiftLeft(b, n).mod(new BigNumber(2).exponentiatedBy(w));
    let y = _dafny.ShiftRight(b, w - n);
    return x.plus(y);
  }
  $module.RotateRight = function(b, n, w) {  // (b >> n) | truncate(b << (w - n))
    let x = _dafny.ShiftRight(b, n);
    let y = _dafny.ShiftLeft(b, w - n).mod(new BigNumber(2).exponentiatedBy(w));;
    return x.plus(y);
  }
  $module.BitwiseAnd = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 & b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    return r;
  }
  $module.BitwiseOr = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 | b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseXor = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 ^ b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseNot = function(a, bits) {
    let r = _dafny.ZERO;
    let h = _dafny.ONE;
    for (let i = 0; i < bits; i++) {
      let bit = a.mod(2);
      if (bit.isZero()) {
        r = r.plus(h);
      }
      a = a.dividedToIntegerBy(2);
      h = h.multipliedBy(2);
    }
    return r;
  }
  $module.Quantifier = function(vals, frall, pred) {
    for (let u of vals) {
      if (pred(u) !== frall) { return !frall; }
    }
    return frall;
  }
  $module.PlusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) + b.charCodeAt(0));
  }
  $module.UnicodePlusChar = function(a, b) {
    return new _dafny.CodePoint(a.value + b.value);
  }
  $module.MinusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) - b.charCodeAt(0));
  }
  $module.UnicodeMinusChar = function(a, b) {
    return new _dafny.CodePoint(a.value - b.value);
  }
  $module.AllBooleans = function*() {
    yield false;
    yield true;
  }
  $module.AllChars = function*() {
    for (let i = 0; i < 0x10000; i++) {
      yield String.fromCharCode(i);
    }
  }
  $module.AllUnicodeChars = function*() {
    for (let i = 0; i < 0xD800; i++) {
      yield new _dafny.CodePoint(i);
    }
    for (let i = 0xE0000; i < 0x110000; i++) {
      yield new _dafny.CodePoint(i);
    }
  }
  $module.AllIntegers = function*() {
    yield _dafny.ZERO;
    for (let j = _dafny.ONE;; j = j.plus(1)) {
      yield j;
      yield j.negated();
    }
  }
  $module.IntegerRange = function*(lo, hi) {
    if (lo === null) {
      while (true) {
        hi = hi.minus(1);
        yield hi;
      }
    } else if (hi === null) {
      while (true) {
        yield lo;
        lo = lo.plus(1);
      }
    } else {
      while (lo.isLessThan(hi)) {
        yield lo;
        lo = lo.plus(1);
      }
    }
  }
  $module.SingleValue = function*(v) {
    yield v;
  }
  $module.HaltException = class HaltException extends Error {
    constructor(message) {
      super(message)
    }
  }
  $module.HandleHaltExceptions = function(f) {
    try {
      f()
    } catch (e) {
      if (e instanceof _dafny.HaltException) {
        process.stdout.write("[Program halted] " + e.message + "\n")
        process.exitCode = 1
      } else {
        throw e
      }
    }
  }
  $module.FromMainArguments = function(args) {
    var a = [...args];
    a.splice(0, 2, args[0] + " " + args[1]);
    return a;
  }
  $module.UnicodeFromMainArguments = function(args) {
    return $module.FromMainArguments(args).map(_dafny.Seq.UnicodeFromString);
  }
  return $module;

  // What follows are routines private to the Dafny runtime
  function buildArray(initValue, ...dims) {
    if (dims.length === 0) {
      return initValue;
    } else {
      let a = Array(dims[0].toNumber());
      let b = Array.from(a, (x) => buildArray(initValue, ...dims.slice(1)));
      return b;
    }
  }
  function arrayElementsToString(a) {
    // like `a.join(", ")`, but calling _dafny.toString(x) on every element x instead of x.toString()
    let s = "";
    let sep = "";
    for (let x of a) {
      s += sep + _dafny.toString(x);
      sep = ", ";
    }
    return s;
  }
  function BigNumberGcd(a, b){  // gcd of two non-negative BigNumber's
    while (true) {
      if (a.isZero()) {
        return b;
      } else if (b.isZero()) {
        return a;
      }
      if (a.isLessThan(b)) {
        b = b.modulo(a);
      } else {
        a = a.modulo(b);
      }
    }
  }
})();
// Dafny program systemModulePopulator.dfy compiled into JavaScript
let _System = (function() {
  let $module = {};

  $module.nat = class nat {
    constructor () {
    }
    static get Default() {
      return _dafny.ZERO;
    }
    static _Is(__source) {
      let _0_x = (__source);
      return (_dafny.ZERO).isLessThanOrEqualTo(_0_x);
    }
  };

  return $module;
})(); // end of module _System
let _module = (function() {
  let $module = {};

  $module.__default = class __default {
    constructor () {
      this._tname = "_module._default";
    }
    _parentTraits() {
      return [];
    }
    static abs(x) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return (new BigNumber(-1)).multipliedBy(x);
      } else {
        return x;
      }
    };
    static safeIndex(x, length) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return _dafny.ZERO;
      } else if ((length).isLessThanOrEqualTo(x)) {
        return (x).mod(length);
      } else {
        return x;
      }
    };
    static safeDivisionInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return _dafny.EuclideanDivision(x1, x2);
      }
    };
    static safeModuloInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return (x1).mod(x2);
      }
    };
    static fm5(globalState) {
      if (!(new BigNumber(608)).isEqualTo(new BigNumber(898))) {
        return new _dafny.CodePoint('v'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('l'.codePointAt(0));
      }
    };
    static fm8(p0, p1, p2, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(100)), function (_0_i0) {
        return (_module.D1.create_DC1(new _dafny.CodePoint('d'.codePointAt(0)))).dtor_cf1;
      });
    };
    static fm9(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,false))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),true));
    };
    static fm10(p0, p1, p2, p3, globalState) {
      let _source0 = _module.D11.create_DC30();
      if (_source0.is_DC29) {
        let _1___mcc_h0 = (_source0).cf43;
        let _2_cf43 = _1___mcc_h0;
        return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-222)), function (_3_i0) {
          return true;
        }), _dafny.Seq.of(false, false, false));
      } else if (_source0.is_DC30) {
        return _dafny.Seq.of(!(false), !(!(!(false))), !(true), true);
      } else if (_source0.is_DC28) {
        let _4___mcc_h1 = (_source0).cf42;
        let _5_cf42 = _4___mcc_h1;
        return _dafny.Seq.Concat(_dafny.Seq.of(true, true, !(!(false)), true), _dafny.Seq.Create(_module.__default.abs(new BigNumber(593)), function (_6_i1) {
          return false;
        }));
      } else {
        let _7___mcc_h2 = (_source0).cf44;
        let _8_cf44 = _7___mcc_h2;
        return _dafny.Seq.Concat(_dafny.Seq.of(!(false)), _dafny.Seq.of(true));
      }
    };
    static fm16(p0, globalState) {
      return (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((((_dafny.Set.fromElements(true, false)).Difference(_dafny.Set.fromElements(false))).Intersect(_dafny.Set.fromElements(!(!(false))))).length)));
    };
    static fm17(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-771)), function (_9_i0) {
        return new _dafny.CodePoint('e'.codePointAt(0));
      })).length)));
    };
    static fm20(p0, p1, p2, globalState) {
      return ((_dafny.ZERO).minus(new BigNumber(-461))).multipliedBy((new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-343)), function (_10_i0) {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('e'.codePointAt(0)),true)).length);
      }))).cardinality())).plus(new BigNumber(518)));
    };
    static fm21(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("yvab"),true)).length)), _dafny.Seq.of(new BigNumber((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(927), new BigNumber(999))) {
          let _11_v0 = _compr_0;
          if (((new BigNumber(927)).isLessThanOrEqualTo(_11_v0)) && ((_11_v0).isLessThan(new BigNumber(999)))) {
            _coll0.push([(_11_v0).plus(new BigNumber((_dafny.Seq.UnicodeFromString("ykwfkuyl")).length)),new BigNumber(298)]);
          }
        }
        return _coll0;
      }()).length), new BigNumber(312), new BigNumber((_dafny.Seq.UnicodeFromString("h")).length), new BigNumber(792))), _dafny.Seq.of(new BigNumber(921)));
    };
    static fm22(globalState) {
      return (_dafny.MultiSet.fromElements(_module.D1.create_DC3(_module.D1.create_DC2()))).Union(_dafny.MultiSet.fromElements(_module.D1.create_DC3(_module.D1.create_DC1(new _dafny.CodePoint('y'.codePointAt(0)))), _module.D1.create_DC3(_module.D1.create_DC2()), _module.D1.create_DC3(_module.D1.create_DC2())));
    };
    static fm23(p0, globalState) {
      let _source1 = false;
      let _12___mcc_h0 = _source1;
      let _13_cf0 = _12___mcc_h0;
      return _module.D1.create_DC3(_module.D1.create_DC2());
    };
    static fm24(p0, p1, p2, globalState) {
      return !(!(new BigNumber((_dafny.Seq.of(!(true), !(false))).length)).isEqualTo(new BigNumber(571)));
    };
    static fm25(globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(false,!(true))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,false)));
    };
    static fm26(p0, p1, globalState) {
      return _module.D3.create_DC6((_dafny.Set.fromElements(new BigNumber(644))).Union(function () {
  let _coll1 = new _dafny.Set();
  for (const _compr_1 of _dafny.IntegerRange(new BigNumber(-907), new BigNumber(-218))) {
    let _14_v0 = _compr_1;
    if (((new BigNumber(-907)).isLessThanOrEqualTo(_14_v0)) && ((_14_v0).isLessThan(new BigNumber(-218)))) {
      _coll1.add((_14_v0).multipliedBy(new BigNumber(-978)));
    }
  }
  return _coll1;
}()));
    };
    static fm27(globalState) {
      return (_dafny.Set.fromElements(_dafny.MultiSet.fromElements(true, true, false))).Intersect(_dafny.Set.fromElements(_dafny.MultiSet.fromElements(true)));
    };
    static fm28(p0, globalState) {
      return (_module.D2.create_DC4(_dafny.MultiSet.FromArray(_dafny.Seq.of(!(true))))).dtor_cf3;
    };
    static fm29(p0, p1, globalState) {
      let _source2 = _module.D10.create_DC25(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(529),new BigNumber(((_module.D2.create_DC4(_dafny.MultiSet.fromElements(false))).dtor_cf3).cardinality())));
      if (_source2.is_DC26) {
        let _15___mcc_h0 = (_source2).cf40;
        let _16_cf40 = _15___mcc_h0;
        return (_module.D3.create_DC6(_dafny.Set.fromElements(new BigNumber((function () {
  let _coll2 = new _dafny.Map();
  for (const _compr_2 of _dafny.IntegerRange(new BigNumber(-914), new BigNumber(258))) {
    let _17_v0 = _compr_2;
    if (((new BigNumber(-914)).isLessThanOrEqualTo(_17_v0)) && ((_17_v0).isLessThan(new BigNumber(258)))) {
      _coll2.push([(_17_v0).minus(_16_cf40),true]);
    }
  }
  return _coll2;
}()).length)))).dtor_cf9;
      } else if (_source2.is_DC25) {
        let _18___mcc_h1 = (_source2).cf39;
        let _19_cf39 = _18___mcc_h1;
        return function () {
          let _coll3 = new _dafny.Set();
          for (const _compr_3 of _dafny.IntegerRange(new BigNumber(574), new BigNumber(360))) {
            let _20_v1 = _compr_3;
            if (((new BigNumber(574)).isLessThanOrEqualTo(_20_v1)) && ((_20_v1).isLessThan(new BigNumber(360)))) {
              _coll3.add(_module.__default.safeDivisionInt(_20_v1, new BigNumber((_dafny.Seq.of(new BigNumber((function () {
                let _coll4 = new _dafny.Set();
                for (const _compr_4 of _dafny.IntegerRange(new BigNumber(772), new BigNumber(533))) {
                  let _21_v2 = _compr_4;
                  if (((new BigNumber(772)).isLessThanOrEqualTo(_21_v2)) && ((_21_v2).isLessThan(new BigNumber(533)))) {
                    _coll4.add((_21_v2).plus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(true, false)).cardinality()), new BigNumber((_dafny.Seq.of(new BigNumber(-308))).length), new BigNumber(-914), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(50))).cardinality()))).length)));
                  }
                }
                return _coll4;
              }()).length))).length)));
            }
          }
          return _coll3;
        }();
      } else {
        let _22___mcc_h2 = (_source2).cf41;
        let _23_cf41 = _22___mcc_h2;
        return _dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("xlnmqckl")).length));
      }
    };
    static fm30(p0, globalState) {
      if ((new BigNumber(720)).isLessThan(new BigNumber((function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of (_dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC6(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("rsyxbepi")).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
  let _coll6 = new _dafny.Set();
  for (const _compr_6 of _dafny.IntegerRange(new BigNumber(853), new BigNumber(631))) {
    let _24_v1 = _compr_6;
    if (((new BigNumber(853)).isLessThanOrEqualTo(_24_v1)) && ((_24_v1).isLessThan(new BigNumber(631)))) {
      _coll6.add((_24_v1).plus(new BigNumber((_dafny.Set.fromElements(false, false)).length)));
    }
  }
  return _coll6;
}()).length),new BigNumber((_dafny.Set.fromElements(true)).length))).length), new BigNumber((_dafny.Set.fromElements(true, false, !(false))).length))).length))),_dafny.Seq.UnicodeFromString("veyog"))).Keys.Elements) {
          let _25_v0 = _compr_5;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC6(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("rsyxbepi")).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
  let _coll7 = new _dafny.Set();
  for (const _compr_7 of _dafny.IntegerRange(new BigNumber(853), new BigNumber(631))) {
    let _26_v1 = _compr_7;
    if (((new BigNumber(853)).isLessThanOrEqualTo(_26_v1)) && ((_26_v1).isLessThan(new BigNumber(631)))) {
      _coll7.add((_26_v1).plus(new BigNumber((_dafny.Set.fromElements(false, false)).length)));
    }
  }
  return _coll7;
}()).length),new BigNumber((_dafny.Set.fromElements(true)).length))).length), new BigNumber((_dafny.Set.fromElements(true, false, !(false))).length))).length))),_dafny.Seq.UnicodeFromString("veyog"))).contains(_25_v0)) {
            _coll5.push([_25_v0,new BigNumber(-649)]);
          }
        }
        return _coll5;
      }()).length))) {
        return _dafny.Seq.Concat(_dafny.Seq.of(!(true)), _dafny.Seq.of(false));
      } else if (!(!(true))) {
        return _dafny.Seq.of(true, true, true, false, true);
      } else {
        return _dafny.Seq.of(true, !(true));
      }
    };
    static fm31(p0, p1, globalState) {
      return _dafny.Set.fromElements((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(!(!(false)), false)).length)),new BigNumber((_dafny.Seq.of(new BigNumber(617))).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(55),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-936)), function (_27_i0) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      })).length))), (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(424),(_module.D8.create_DC22(new BigNumber((_dafny.Seq.of(new BigNumber(-954), new BigNumber((_dafny.Seq.UnicodeFromString("qejrjkfg")).length))).length))).dtor_cf35)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(140)), function (_28_i1) {
        return new BigNumber(360);
      })).length),new BigNumber((_dafny.Seq.UnicodeFromString("dhvkev")).length))), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(713),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(108)), function (_29_i2) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      })).length)));
    };
    static fm32(globalState) {
      let _source3 = _module.D5.create_DC13(false, new _dafny.CodePoint('y'.codePointAt(0)));
      if (_source3.is_DC12) {
        let _30___mcc_h0 = (_source3).cf19;
        let _31_cf19 = _30___mcc_h0;
        return _module.D5.create_DC13(_31_cf19, new _dafny.CodePoint('r'.codePointAt(0)));
      } else if (_source3.is_DC13) {
        let _32___mcc_h1 = (_source3).cf20;
        let _33___mcc_h2 = (_source3).cf21;
        let _34_cf21 = _33___mcc_h2;
        let _35_cf20 = _32___mcc_h1;
        return _module.D5.create_DC13(_35_cf20, _34_cf21);
      } else if (_source3.is_DC11) {
        let _36___mcc_h3 = (_source3).cf18;
        let _37_cf18 = _36___mcc_h3;
        return _module.D5.create_DC13(true, new _dafny.CodePoint('f'.codePointAt(0)));
      } else {
        let _38___mcc_h4 = (_source3).cf22;
        let _39_cf22 = _38___mcc_h4;
        return _module.D5.create_DC13(false, new _dafny.CodePoint('f'.codePointAt(0)));
      }
    };
    static fm33(globalState) {
      return ((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("qqnloowd")).length)))).isEqualTo(new BigNumber(-548));
    };
    static fm34(p0, p1, globalState) {
      return (_dafny.Set.fromElements(_module.D6.create_DC16(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(false),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(155))).length)), _dafny.Seq.of(true)), _module.D6.create_DC16(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(false),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true)).length))).length)), _dafny.Seq.of(false)))).Intersect(function () {
        let _coll8 = new _dafny.Set();
        for (const _compr_8 of (_dafny.MultiSet.FromArray(_dafny.Seq.of(_module.D6.create_DC16(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(true),new BigNumber(699)), _dafny.Seq.of(false, true))))).Elements) {
          let _40_v0 = _compr_8;
          if ((_dafny.MultiSet.FromArray(_dafny.Seq.of(_module.D6.create_DC16(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(true),new BigNumber(699)), _dafny.Seq.of(false, true))))).contains(_40_v0)) {
            _coll8.add(_40_v0);
          }
        }
        return _coll8;
      }());
    };
    static fm35(p0, globalState) {
      return _module.D6.create_DC16((function () {
  let _coll9 = new _dafny.Map();
  for (const _compr_9 of (_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(true), _dafny.MultiSet.fromElements(true))).Elements) {
    let _41_v0 = _compr_9;
    if ((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(true), _dafny.MultiSet.fromElements(true))).contains(_41_v0)) {
      _coll9.push([_41_v0,new BigNumber((_dafny.Seq.of(false)).length)]);
    }
  }
  return _coll9;
}()).Merge(function () {
  let _coll10 = new _dafny.Map();
  for (const _compr_10 of (function () {
    let _coll11 = new _dafny.Map();
    for (const _compr_11 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(83)), function (_42_i0) {
      return _dafny.MultiSet.fromElements(false, true);
    })).Elements) {
      let _43_v2 = _compr_11;
      if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(83)), function (_42_i0) {
        return _dafny.MultiSet.fromElements(false, true);
      }), _43_v2)) {
        _coll11.push([_43_v2,new BigNumber(775)]);
      }
    }
    return _coll11;
  }()).Keys.Elements) {
    let _44_v1 = _compr_10;
    if ((function () {
      let _coll12 = new _dafny.Map();
      for (const _compr_12 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(83)), function (_42_i0) {
        return _dafny.MultiSet.fromElements(false, true);
      })).Elements) {
        let _43_v2 = _compr_12;
        if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(83)), function (_42_i0) {
          return _dafny.MultiSet.fromElements(false, true);
        }), _43_v2)) {
          _coll12.push([_43_v2,new BigNumber(775)]);
        }
      }
      return _coll12;
    }()).contains(_44_v1)) {
      _coll10.push([_44_v1,new BigNumber((_dafny.Seq.of(true, true, false, false, !(false))).length)]);
    }
  }
  return _coll10;
}()), _dafny.Seq.of(false));
    };
    static fm36(p0, p1, globalState) {
      return (function () {
        let _coll13 = new _dafny.Set();
        for (const _compr_13 of (_dafny.Set.fromElements(_dafny.MultiSet.fromElements(_module.D1.create_DC3(_module.D1.create_DC2()), _module.D1.create_DC3(_module.D1.create_DC1(new _dafny.CodePoint('f'.codePointAt(0))))))).Elements) {
          let _45_v0 = _compr_13;
          if ((_dafny.Set.fromElements(_dafny.MultiSet.fromElements(_module.D1.create_DC3(_module.D1.create_DC2()), _module.D1.create_DC3(_module.D1.create_DC1(new _dafny.CodePoint('f'.codePointAt(0))))))).contains(_45_v0)) {
            _coll13.add(_45_v0);
          }
        }
        return _coll13;
      }()).Union((_dafny.Set.fromElements(_dafny.MultiSet.fromElements(_module.D1.create_DC3(_module.D1.create_DC3(_module.D1.create_DC2())), _module.D1.create_DC3(_module.D1.create_DC2()), _module.D1.create_DC3(_module.D1.create_DC2()), _module.D1.create_DC3(_module.D1.create_DC2())), _dafny.MultiSet.fromElements(_module.D1.create_DC3(_module.D1.create_DC3(_module.D1.create_DC2()))))).Intersect(_dafny.Set.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.of(_module.D1.create_DC3(_module.D1.create_DC2()), _module.D1.create_DC3(_module.D1.create_DC1(new _dafny.CodePoint('m'.codePointAt(0)))), _module.D1.create_DC3(_module.D1.create_DC2()))))));
    };
    static fm37(globalState) {
      if (!((new BigNumber((_dafny.Seq.UnicodeFromString("npm")).length)).isLessThan(new BigNumber(-818)))) {
        return function () {
          let _coll14 = new _dafny.Map();
          for (const _compr_14 of (_dafny.Set.fromElements(new BigNumber(93), new BigNumber((_dafny.Seq.of(false)).length), new BigNumber(694), new BigNumber(980))).Elements) {
            let _46_v0 = _compr_14;
            if ((_dafny.Set.fromElements(new BigNumber(93), new BigNumber((_dafny.Seq.of(false)).length), new BigNumber(694), new BigNumber(980))).contains(_46_v0)) {
              _coll14.push([(_46_v0).plus(new BigNumber(302)),_dafny.Seq.Create(_module.__default.abs(new BigNumber(184)), function (_47_i0) {
                return new _dafny.CodePoint('f'.codePointAt(0));
              })]);
            }
          }
          return _coll14;
        }();
      } else {
        return (function () {
          let _coll15 = new _dafny.Map();
          for (const _compr_15 of (_dafny.Seq.of(new BigNumber(732), new BigNumber(894))).Elements) {
            let _48_v1 = _compr_15;
            if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(732), new BigNumber(894)), _48_v1)) {
              _coll15.push([(_48_v1).minus(new BigNumber(-626)),_dafny.Seq.UnicodeFromString("nf")]);
            }
          }
          return _coll15;
        }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("r")).length),_dafny.Seq.Create(_module.__default.abs(new BigNumber(-552)), function (_49_i1) {
          return new _dafny.CodePoint('r'.codePointAt(0));
        })));
      }
    };
    static fm38(p0, p1, p2, p3, globalState) {
      if ((new BigNumber((function () {
        let _coll16 = new _dafny.Map();
        for (const _compr_16 of (_dafny.Seq.of(_dafny.MultiSet.fromElements(false), _dafny.MultiSet.FromArray(_dafny.Seq.of(false, false)))).Elements) {
          let _50_v0 = _compr_16;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.MultiSet.fromElements(false), _dafny.MultiSet.FromArray(_dafny.Seq.of(false, false))), _50_v0)) {
            _coll16.push([_50_v0,new BigNumber(-237)]);
          }
        }
        return _coll16;
      }()).length)).isLessThan(new BigNumber(47))) {
        return (_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-275)),new BigNumber(295))).Merge(function () {
          let _coll17 = new _dafny.Map();
          for (const _compr_17 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(!(false))).cardinality()),false)).Keys.Elements) {
            let _51_v1 = _compr_17;
            if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(!(false))).cardinality()),false)).contains(_51_v1)) {
              _coll17.push([(_51_v1).multipliedBy((_dafny.ZERO).minus(new BigNumber(-744))),new BigNumber(-645)]);
            }
          }
          return _coll17;
        }());
      } else {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
          let _coll18 = new _dafny.Map();
          for (const _compr_18 of _dafny.IntegerRange(new BigNumber(887), new BigNumber(-826))) {
            let _52_v2 = _compr_18;
            if (((new BigNumber(887)).isLessThanOrEqualTo(_52_v2)) && ((_52_v2).isLessThan(new BigNumber(-826)))) {
              _coll18.push([_module.__default.safeModuloInt(_52_v2, new BigNumber(-828)),true]);
            }
          }
          return _coll18;
        }()).length),new BigNumber(28));
      }
    };
    static fm39(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((function () {
        let _coll19 = new _dafny.Map();
        for (const _compr_19 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(500),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(!(true))).cardinality())),true)).length))).Keys.Elements) {
          let _53_v0 = _compr_19;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(500),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(!(true))).cardinality())),true)).length))).contains(_53_v0)) {
            _coll19.push([_module.__default.safeModuloInt(_53_v0, new BigNumber(-296)),_dafny.Set.fromElements(true)]);
          }
        }
        return _coll19;
      }()).length)))).cardinality()))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Set.fromElements(true, true)).length)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
        let _coll20 = new _dafny.Map();
        for (const _compr_20 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(667)), function (_54_i0) {
          return new BigNumber(405);
        })).Elements) {
          let _55_v1 = _compr_20;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(667)), function (_54_i0) {
            return new BigNumber(405);
          }), _55_v1)) {
            _coll20.push([(_55_v1).plus(new BigNumber((_dafny.Seq.UnicodeFromString("qqnmhka")).length)),new BigNumber(965)]);
          }
        }
        return _coll20;
      }()).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(!(!(true))),new BigNumber(224))));
    };
    static m0(p0, p1, p2, globalState) {
      let r0 = [];
      let r1 = false;
      let _56_v0;
      _56_v0 = _dafny.Seq.of(new BigNumber(-104));
      let _57_v1;
      _57_v1 = _dafny.Map.Empty.slice().updateUnsafe(_56_v0,p2);
      let _58_v2;
      _58_v2 = _module.D7.create_DC18(p2, (((_57_v1).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(748)), ((_61_p0) => function (_62_i0) {
  return _61_p0;
})(p0)))) ? ((_57_v1).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(748)), ((_59_p0) => function (_60_i0) {
  return _59_p0;
})(p0)))) : (p1)));
      let _source4 = _58_v2;
      if (_source4.is_DC18) {
        let _63___mcc_h0 = (_source4).cf27;
        let _64___mcc_h1 = (_source4).cf28;
        let _65_cf28 = _64___mcc_h1;
        let _66_cf27 = _63___mcc_h0;
        let _67_v3;
        _67_v3 = _dafny.Seq.of(!(true));
        let _68_v4;
        _68_v4 = _dafny.Map.Empty.slice().updateUnsafe(_module.D7.create_DC17(_dafny.Seq.UnicodeFromString("mlycpc")),_67_v3);
        let _69_v5;
        _69_v5 = _dafny.Map.Empty.slice().updateUnsafe(_68_v4,((p2) ? (true) : (p1)));
        _69_v5 = (_69_v5).update(_68_v4, _65_cf28);
        let _70_v6;
        _70_v6 = _dafny.Map.Empty.slice().updateUnsafe(p0,_66_cf27);
        let _71_v7;
        _71_v7 = _dafny.Map.Empty.slice().updateUnsafe(_70_v6,p0);
        let _72_v8;
        _72_v8 = _dafny.Seq.UnicodeFromString("n");
        let _73_v9;
        let _nw0 = new _module.C5();
        _nw0.__ctor(_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(911)), function (_74_i1) {
          return new _dafny.CodePoint('k'.codePointAt(0));
        }), _module.__default.safeIndex(new BigNumber((_71_v7).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(911)), function (_75_i1) {
          return new _dafny.CodePoint('k'.codePointAt(0));
        })).length)), new _dafny.CodePoint('n'.codePointAt(0))), _72_v8), (_dafny.ZERO).minus((_dafny.ZERO).minus(p0)), p0, p0);
        _73_v9 = _nw0;
        let _76_v11;
        let _nw1 = new _module.C4();
        _nw1.__ctor(p1, new BigNumber((function () {
          let _coll21 = new _dafny.Map();
          for (const _compr_21 of (_70_v6).Keys.Elements) {
            let _77_v10 = _compr_21;
            if ((_70_v6).contains(_77_v10)) {
              _coll21.push([_module.__default.safeModuloInt(_77_v10, (_73_v9).f10),(_73_v9).f10]);
            }
          }
          return _coll21;
        }()).length), new BigNumber((_module.__default.fm21(p0, (_dafny.ZERO).minus(p0), p0, _dafny.Seq.UnicodeFromString("etffgh"), globalState)).length));
        _76_v11 = _nw1;
        let _78_v12;
        _78_v12 = _dafny.MultiSet.fromElements(_76_v11);
        let _79_v13;
        _79_v13 = _dafny.Seq.of(_76_v11);
        let _80_v14;
        _80_v14 = _dafny.MultiSet.fromElements(_78_v12, _78_v12, _dafny.MultiSet.fromElements(_76_v11, _76_v11), _dafny.MultiSet.FromArray(_79_v13));
        (globalState).f3 = (_dafny.MultiSet.fromElements(_78_v12, _dafny.MultiSet.fromElements(_76_v11))).IsSubsetOf(_80_v14);
        let _81_v15;
        _81_v15 = new BigNumber(-986);
        _81_v15 = new BigNumber((_dafny.Set.fromElements((_73_v9).f10, new BigNumber(877), p0, _81_v15)).length);
      } else if (_source4.is_DC19) {
        let _82___mcc_h2 = (_source4).cf29;
        let _83___mcc_h3 = (_source4).cf30;
        let _84___mcc_h4 = (_source4).cf31;
        let _85_cf31 = _84___mcc_h4;
        let _86_cf30 = _83___mcc_h3;
        let _87_cf29 = _82___mcc_h2;
        let _88_v16;
        let _nw2 = new _module.C3();
        _nw2.__ctor(p2, _module.__default.safeModuloInt(_module.__default.fm20(p2, p1, !(p1), globalState), _85_cf31), new BigNumber(-709));
        _88_v16 = _nw2;
        let _89_v17;
        _89_v17 = _dafny.Seq.UnicodeFromString("rtmdkqwc");
        let _90_v18;
        _90_v18 = _dafny.Set.fromElements((_88_v16).fm13(_89_v17, globalState), p0);
        let _91_v19;
        _91_v19 = _dafny.Map.Empty.slice().updateUnsafe((_90_v18).Intersect(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(p1, !(!(!(p1))))).length), p0)),(((_88_v16).f12) ? (true) : (_86_cf30)));
        _91_v19 = (_91_v19).update(_90_v18, !(_86_cf30) || ((_88_v16).f12));
        _85_cf31 = _85_cf31;
        _86_cf30 = _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("ruevlxjng"), _dafny.Seq.Concat(_89_v17, _89_v17));
      } else if (_source4.is_DC20) {
        let _92___mcc_h5 = (_source4).cf32;
        let _93___mcc_h6 = (_source4).cf33;
        let _94_cf33 = _93___mcc_h6;
        let _95_cf32 = _92___mcc_h5;
        (globalState).f4 = p1;
        let _96_v20;
        let _nw3 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
        _96_v20 = _nw3;
        let _97_v21;
        let _nw4 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Map.Empty);
        _97_v21 = _nw4;
        let _98_v22;
        _98_v22 = _dafny.Seq.UnicodeFromString("xjka");
        let _99_v23;
        _99_v23 = _dafny.Map.Empty.slice().updateUnsafe(_94_cf33,_98_v22);
        let _index0 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_97_v21).length));
        (_97_v21)[_index0] = _99_v23;
        let _100_v24;
        _100_v24 = _module.D9.create_DC23(_96_v20);
        let _101_v25;
        _101_v25 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
        let _102_v26;
        _102_v26 = _dafny.Seq.of(p1, !(p1) || ((((_101_v25).contains(false)) ? ((_101_v25).get(false)) : (p1))));
        let _103_v27;
        _103_v27 = _dafny.MultiSet.fromElements(new BigNumber(-877));
        let _104_v28;
        _104_v28 = _dafny.Set.fromElements(p0);
        let _105_v29;
        let _nw5 = new _module.C5();
        _nw5.__ctor(_98_v22, _94_cf33, new BigNumber((_103_v27).cardinality()), new BigNumber((_104_v28).length));
        _105_v29 = _nw5;
        let _106_v30;
        _106_v30 = _dafny.Map.Empty.slice().updateUnsafe(_105_v29,p2);
        let _107_v31;
        _107_v31 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_105_v29).f10, p0),(_module.__default.fm37(globalState)).Merge(_module.__default.fm37(globalState)));
        let _index1 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_97_v21).length));
        let _rhs0 = p1;
        let _rhs1 = (_100_v24).dtor_cf36;
        let _rhs2 = (_102_v26)[_module.__default.safeIndex((new BigNumber((_106_v30).length)).multipliedBy((_105_v29).f10), new BigNumber((_102_v26).length))];
        let _rhs3 = (((_107_v31).contains((_dafny.ZERO).minus((_105_v29).f10))) ? ((_107_v31).get((_dafny.ZERO).minus((_105_v29).f10))) : ((_99_v23).Merge(_99_v23)));
        let _lhs0 = globalState;
        let _lhs1 = globalState;
        let _lhs2 = _97_v21;
        let _lhs3 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_97_v21).length));
        _lhs0.f4 = _rhs0;
        _96_v20 = _rhs1;
        _lhs1.f4 = _rhs2;
        _lhs2[_lhs3] = _rhs3;
        let _108_v32;
        let _nw6 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _108_v32 = _nw6;
        let _109_v33;
        _109_v33 = _dafny.Map.Empty.slice().updateUnsafe(_108_v32,new BigNumber(-969));
        _109_v33 = (_109_v33).update(_108_v32, _94_cf33);
        _94_cf33 = _module.__default.safeModuloInt((_94_cf33).minus(_94_cf33), p0);
      } else {
        let _110___mcc_h7 = (_source4).cf26;
        let _111_cf26 = _110___mcc_h7;
        let _112_v34;
        _112_v34 = _dafny.Seq.of(!(true));
        let _113_v35;
        let _nw7 = Array((new BigNumber(18)).toNumber());
        _nw7[(_dafny.ZERO).toNumber()] = new BigNumber((_module.__default.fm30(false, globalState)).length);
        _nw7[(_dafny.ONE).toNumber()] = new BigNumber(122);
        _nw7[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.Seq.of(true)).length);
        _nw7[(new BigNumber(3)).toNumber()] = _module.__default.fm20(p1, p2, p1, globalState);
        _nw7[(new BigNumber(4)).toNumber()] = p0;
        _nw7[(new BigNumber(5)).toNumber()] = p0;
        _nw7[(new BigNumber(6)).toNumber()] = p0;
        _nw7[(new BigNumber(7)).toNumber()] = p0;
        _nw7[(new BigNumber(8)).toNumber()] = p0;
        _nw7[(new BigNumber(9)).toNumber()] = p0;
        _nw7[(new BigNumber(10)).toNumber()] = new BigNumber(674);
        _nw7[(new BigNumber(11)).toNumber()] = p0;
        _nw7[(new BigNumber(12)).toNumber()] = p0;
        _nw7[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_112_v34).length));
        _nw7[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus(p0);
        _nw7[(new BigNumber(15)).toNumber()] = p0;
        _nw7[(new BigNumber(16)).toNumber()] = p0;
        _nw7[(new BigNumber(17)).toNumber()] = new BigNumber(885);
        _113_v35 = _nw7;
        let _114_v36;
        _114_v36 = _dafny.Seq.of(_113_v35, _113_v35);
        let _115_v37;
        _115_v37 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_114_v36)[_module.__default.safeIndex(p0, new BigNumber((_114_v36).length))]);
        let _116_v38;
        _116_v38 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),p1);
        _115_v37 = (_115_v37).update((((_116_v38).contains(p0)) ? ((_116_v38).get(p0)) : (p1)), _113_v35);
        (globalState).f4 = ((p2) ? (((p2) ? (false) : (p1))) : (true));
        if ((new BigNumber(240)).isLessThan(p0)) {
          let _117_v39;
          let _nw8 = new _module.C4();
          _nw8.__ctor(p1, new BigNumber((_dafny.Seq.update(_module.__default.fm30(!(!(false)), globalState), _module.__default.safeIndex(p0, new BigNumber((_module.__default.fm30(!(!(false)), globalState)).length)), p1)).length), (_dafny.ZERO).minus(p0));
          _117_v39 = _nw8;
          let _118_v40;
          let _nw9 = new _module.C1();
          _nw9.__ctor(p0, new BigNumber((_dafny.Seq.UnicodeFromString("mhnx")).length));
          _118_v40 = _nw9;
          let _119_v41;
          _119_v41 = _dafny.Seq.of(_118_v40, _118_v40);
          let _rhs4 = _117_v39;
          let _rhs5 = _119_v41;
          _117_v39 = _rhs4;
          _119_v41 = _rhs5;
          let _120_v42;
          _120_v42 = _dafny.ZERO;
          _120_v42 = _module.__default.safeModuloInt(_module.__default.safeModuloInt((_118_v40).f8, (_118_v40).f7), (_dafny.ZERO).minus((_dafny.ZERO).minus(_120_v42)));
          _56_v0 = ((p1) ? (_56_v0) : (_56_v0));
          let _index2 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_113_v35).length));
          (_113_v35)[_index2] = (_dafny.ZERO).minus(p0);
          let _121_v43;
          _121_v43 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
          let _index3 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_113_v35).length));
          (_113_v35)[_index3] = (_120_v42).multipliedBy((((_121_v43).contains(_module.__default.fm24((_117_v39).f11, (_118_v40).f7, _111_cf26, globalState))) ? ((_121_v43).get(_module.__default.fm24((_117_v39).f11, (_118_v40).f7, _111_cf26, globalState))) : (_120_v42)));
          let _122_v44;
          _122_v44 = _dafny.MultiSet.fromElements(new BigNumber((_module.__default.fm28(p2, globalState)).cardinality()), _120_v42, new BigNumber((_111_cf26).length), _120_v42);
          let _123_v45;
          let _nw10 = Array((new BigNumber(9)).toNumber());
          _nw10[(_dafny.ZERO).toNumber()] = (_117_v39).f11;
          _nw10[(_dafny.ONE).toNumber()] = p2;
          _nw10[(new BigNumber(2)).toNumber()] = (_120_v42).isLessThanOrEqualTo((_118_v40).f7);
          _nw10[(new BigNumber(3)).toNumber()] = p2;
          _nw10[(new BigNumber(4)).toNumber()] = !(p1) || (p1);
          _nw10[(new BigNumber(5)).toNumber()] = p2;
          _nw10[(new BigNumber(6)).toNumber()] = true;
          _nw10[(new BigNumber(7)).toNumber()] = (_122_v44).equals(_122_v44);
          _nw10[(new BigNumber(8)).toNumber()] = (_117_v39).f11;
          _123_v45 = _nw10;
          let _124_v46;
          _124_v46 = _module.D10.create_DC25(_module.__default.fm38((_117_v39).fm1(_dafny.Set.fromElements(p2), p0, (_dafny.ZERO).minus(_120_v42), new BigNumber(554), globalState), true, _module.__default.fm24(p2, new BigNumber(816), _111_cf26, globalState), (_118_v40).f8, globalState));
          let _125_v47;
          _125_v47 = _dafny.Set.fromElements(new BigNumber(((_124_v46).dtor_cf39).length));
          let _index4 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_123_v45).length));
          (_123_v45)[_index4] = (_125_v47).IsSubsetOf(function () {
            let _coll22 = new _dafny.Set();
            for (const _compr_22 of _dafny.IntegerRange(new BigNumber(107), new BigNumber(642))) {
              let _126_v48 = _compr_22;
              if (((new BigNumber(107)).isLessThanOrEqualTo(_126_v48)) && ((_126_v48).isLessThan(new BigNumber(642)))) {
                _coll22.add((_126_v48).multipliedBy((_118_v40).f8));
              }
            }
            return _coll22;
          }());
          let _127_v49;
          let _nw11 = Array((new BigNumber(29)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _127_v49 = _nw11;
          let _index5 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_127_v49).length));
          (_127_v49)[_index5] = new _dafny.CodePoint('e'.codePointAt(0));
          let _index6 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_123_v45).length));
          (_123_v45)[_index6] = p1;
          let _128_v50;
          let _init0 = ((_129_p1, _130_v40) => function (_131_i2) {
            return _dafny.Map.Empty.slice().updateUnsafe(_129_p1,(_130_v40).f8);
          })(p1, _118_v40);
          let _nw12 = Array((new BigNumber(18)).toNumber());
          for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw12.length); _i0_0++) {
            _nw12[_i0_0] = _init0(new BigNumber(_i0_0));
          }
          _128_v50 = _nw12;
          let _index7 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_128_v50).length));
          (_128_v50)[_index7] = (_121_v43).Merge(_121_v43);
          let _132_v51;
          _132_v51 = new _dafny.CodePoint('v'.codePointAt(0));
          let _133_v52;
          _133_v52 = _module.D3.create_DC6(_125_v47);
          let _index8 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_123_v45).length));
          let _index9 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_127_v49).length));
          let _index10 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_123_v45).length));
          let _index11 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_128_v50).length));
          let _rhs6 = (_112_v34)[_module.__default.safeIndex((_118_v40).f8, new BigNumber((_112_v34).length))];
          let _rhs7 = _120_v42;
          let _rhs8 = _132_v51;
          let _rhs9 = p2;
          let _rhs10 = _module.__default.fm39((_118_v40).f8, _133_v52, p2, (_117_v39).f11, globalState);
          let _lhs4 = _123_v45;
          let _lhs5 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_123_v45).length));
          let _lhs6 = _127_v49;
          let _lhs7 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_127_v49).length));
          let _lhs8 = _123_v45;
          let _lhs9 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_123_v45).length));
          let _lhs10 = _128_v50;
          let _lhs11 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_128_v50).length));
          _lhs4[_lhs5] = _rhs6;
          _120_v42 = _rhs7;
          _lhs6[_lhs7] = _rhs8;
          _lhs8[_lhs9] = _rhs9;
          _lhs10[_lhs11] = _rhs10;
        } else {
          let _index12 = _module.__default.safeIndex(new BigNumber(18), new BigNumber((_113_v35).length));
          (_113_v35)[_index12] = p0;
          let _index13 = _module.__default.safeIndex(new BigNumber(18), new BigNumber((_113_v35).length));
          (_113_v35)[_index13] = _module.__default.safeModuloInt(new BigNumber((_56_v0).length), p0);
          let _index14 = _module.__default.safeIndex(new BigNumber(18), new BigNumber((_113_v35).length));
          (_113_v35)[_index14] = (_113_v35)[_module.__default.safeIndex(new BigNumber(18), new BigNumber((_113_v35).length))];
          let _index15 = _module.__default.safeIndex(new BigNumber(18), new BigNumber((_113_v35).length));
          (_113_v35)[_index15] = p0;
          let _134_v53;
          _134_v53 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
          let _135_v54;
          _135_v54 = new _dafny.CodePoint('c'.codePointAt(0));
          let _136_v55;
          let _nw13 = Array((new BigNumber(23)).toNumber());
          _nw13[(_dafny.ZERO).toNumber()] = p2;
          _nw13[(_dafny.ONE).toNumber()] = p2;
          _nw13[(new BigNumber(2)).toNumber()] = p2;
          _nw13[(new BigNumber(3)).toNumber()] = !(_module.__default.fm24(p2, new BigNumber((_134_v53).length), _dafny.Seq.of(_135_v54), globalState));
          _nw13[(new BigNumber(4)).toNumber()] = !(p1);
          _nw13[(new BigNumber(5)).toNumber()] = p2;
          _nw13[(new BigNumber(6)).toNumber()] = p1;
          _nw13[(new BigNumber(7)).toNumber()] = p2;
          _nw13[(new BigNumber(8)).toNumber()] = p2;
          _nw13[(new BigNumber(9)).toNumber()] = p2;
          _nw13[(new BigNumber(10)).toNumber()] = p1;
          _nw13[(new BigNumber(11)).toNumber()] = p2;
          _nw13[(new BigNumber(12)).toNumber()] = p2;
          _nw13[(new BigNumber(13)).toNumber()] = p2;
          _nw13[(new BigNumber(14)).toNumber()] = p1;
          _nw13[(new BigNumber(15)).toNumber()] = p1;
          _nw13[(new BigNumber(16)).toNumber()] = true;
          _nw13[(new BigNumber(17)).toNumber()] = p1;
          _nw13[(new BigNumber(18)).toNumber()] = p1;
          _nw13[(new BigNumber(19)).toNumber()] = p1;
          _nw13[(new BigNumber(20)).toNumber()] = p2;
          _nw13[(new BigNumber(21)).toNumber()] = p1;
          _nw13[(new BigNumber(22)).toNumber()] = true;
          _136_v55 = _nw13;
          let _137_v56;
          _137_v56 = _dafny.Seq.of(_136_v55);
          let _138_v57;
          _138_v57 = _dafny.MultiSet.fromElements((_113_v35)[_module.__default.safeIndex(new BigNumber(18), new BigNumber((_113_v35).length))], new BigNumber((_111_cf26).length));
          let _index16 = _module.__default.safeIndex(new BigNumber(18), new BigNumber((_113_v35).length));
          (_113_v35)[_index16] = ((_module.D4.create_DC9(p0, _137_v56, (_113_v35)[_module.__default.safeIndex(new BigNumber(18), new BigNumber((_113_v35).length))], (((_138_v57).contains((((_138_v57).contains(new BigNumber(516))) ? ((_138_v57).get(new BigNumber(516))) : ((_113_v35)[_module.__default.safeIndex(new BigNumber(18), new BigNumber((_113_v35).length))])))) ? ((_138_v57).get((((_138_v57).contains(new BigNumber(516))) ? ((_138_v57).get(new BigNumber(516))) : ((_113_v35)[_module.__default.safeIndex(new BigNumber(18), new BigNumber((_113_v35).length))])))) : (new BigNumber(584))))).dtor_cf15).plus((_dafny.ZERO).minus((_113_v35)[_module.__default.safeIndex(new BigNumber(18), new BigNumber((_113_v35).length))]));
          let _139_v58;
          let _nw14 = Array((new BigNumber(16)).toNumber()).fill([]);
          _139_v58 = _nw14;
          let _140_v59;
          let _nw15 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
          _140_v59 = _nw15;
          let _index17 = _module.__default.safeIndex(new BigNumber(22), new BigNumber((_139_v58).length));
          (_139_v58)[_index17] = _140_v59;
          let _index18 = _module.__default.safeIndex(new BigNumber(22), new BigNumber((_139_v58).length));
          (_139_v58)[_index18] = _140_v59;
        }
        let _141_v60;
        let _nw16 = new _module.C3();
        _nw16.__ctor(p1, p0, p0);
        _141_v60 = _nw16;
        let _142_v61;
        _142_v61 = _module.D8.create_DC21(_141_v60);
        let _143_v62;
        _143_v62 = _dafny.Map.Empty.slice().updateUnsafe(_142_v61,p2);
        _143_v62 = (_143_v62).update(_module.D8.create_DC21(_141_v60), false);
      }
      let _144_v63;
      _144_v63 = new _dafny.CodePoint('s'.codePointAt(0));
      let _145_v64;
      let _nw17 = Array((new BigNumber(9)).toNumber());
      _nw17[(_dafny.ZERO).toNumber()] = !(!(p2));
      _nw17[(_dafny.ONE).toNumber()] = p2;
      _nw17[(new BigNumber(2)).toNumber()] = p1;
      _nw17[(new BigNumber(3)).toNumber()] = p2;
      _nw17[(new BigNumber(4)).toNumber()] = p1;
      _nw17[(new BigNumber(5)).toNumber()] = true;
      _nw17[(new BigNumber(6)).toNumber()] = p1;
      _nw17[(new BigNumber(7)).toNumber()] = false;
      _nw17[(new BigNumber(8)).toNumber()] = p1;
      _145_v64 = _nw17;
      let _146_v65;
      _146_v65 = _module.D6.create_DC15(_145_v64);
      let _147_v66;
      _147_v66 = _dafny.Seq.UnicodeFromString("nbthqcwcx");
      let _148_v67;
      _148_v67 = _dafny.Map.Empty.slice().updateUnsafe(_146_v65,_147_v66);
      let _149_i3;
      _149_i3 = _dafny.ZERO;
      L0: {
        while (_dafny.Seq.contains((((_148_v67).contains(_146_v65)) ? ((_148_v67).get(_146_v65)) : (_147_v66)), _144_v63)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_149_i3)) {
              break L0;
            }
            _149_i3 = (_149_i3).plus(_dafny.ONE);
            let _150_v68;
            let _nw18 = new _module.C5();
            _nw18.__ctor(_dafny.Seq.UnicodeFromString("e"), p0, p0, p0);
            _150_v68 = _nw18;
            let _151_v69;
            let _nw19 = new _module.C5();
            _nw19.__ctor(_147_v66, (_150_v68).f8, new BigNumber((_dafny.Seq.UnicodeFromString("ks")).length), (_150_v68).f7);
            _151_v69 = _nw19;
            let _152_v70;
            _152_v70 = _module.D11.create_DC28(_151_v69);
            let _153_v71;
            let _nw20 = Array((new BigNumber(15)).toNumber());
            _nw20[(_dafny.ZERO).toNumber()] = _151_v69;
            _nw20[(_dafny.ONE).toNumber()] = _151_v69;
            _nw20[(new BigNumber(2)).toNumber()] = _151_v69;
            _nw20[(new BigNumber(3)).toNumber()] = (_152_v70).dtor_cf42;
            _nw20[(new BigNumber(4)).toNumber()] = _151_v69;
            _nw20[(new BigNumber(5)).toNumber()] = _151_v69;
            _nw20[(new BigNumber(6)).toNumber()] = _151_v69;
            _nw20[(new BigNumber(7)).toNumber()] = _151_v69;
            _nw20[(new BigNumber(8)).toNumber()] = _151_v69;
            _nw20[(new BigNumber(9)).toNumber()] = _151_v69;
            _nw20[(new BigNumber(10)).toNumber()] = _151_v69;
            _nw20[(new BigNumber(11)).toNumber()] = _151_v69;
            _nw20[(new BigNumber(12)).toNumber()] = _151_v69;
            _nw20[(new BigNumber(13)).toNumber()] = _151_v69;
            _nw20[(new BigNumber(14)).toNumber()] = _151_v69;
            _153_v71 = _nw20;
            let _154_v72;
            _154_v72 = _dafny.Map.Empty.slice().updateUnsafe(_150_v68,_153_v71);
            let _155_v73;
            _155_v73 = _module.D12.create_DC32(_154_v72);
            let _156_v74;
            _156_v74 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_150_v68,_153_v71));
            let _157_v75;
            _157_v75 = _dafny.Map.Empty.slice().updateUnsafe((_150_v68).f8,_154_v72);
            let _158_v76;
            let _nw21 = Array((new BigNumber(21)).toNumber());
            _nw21[(_dafny.ZERO).toNumber()] = _154_v72;
            _nw21[(_dafny.ONE).toNumber()] = ((_155_v73).dtor_cf45).Merge(_154_v72);
            _nw21[(new BigNumber(2)).toNumber()] = (_156_v74)[_module.__default.safeIndex(new BigNumber(709), new BigNumber((_156_v74).length))];
            _nw21[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_150_v68,_153_v71);
            _nw21[(new BigNumber(4)).toNumber()] = _154_v72;
            _nw21[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_150_v68,_153_v71);
            _nw21[(new BigNumber(6)).toNumber()] = _154_v72;
            _nw21[(new BigNumber(7)).toNumber()] = _154_v72;
            _nw21[(new BigNumber(8)).toNumber()] = _154_v72;
            _nw21[(new BigNumber(9)).toNumber()] = (_154_v72).update(_150_v68, _153_v71);
            _nw21[(new BigNumber(10)).toNumber()] = _154_v72;
            _nw21[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_150_v68,_153_v71);
            _nw21[(new BigNumber(12)).toNumber()] = _154_v72;
            _nw21[(new BigNumber(13)).toNumber()] = _154_v72;
            _nw21[(new BigNumber(14)).toNumber()] = (_154_v72).Merge(_154_v72);
            _nw21[(new BigNumber(15)).toNumber()] = _154_v72;
            _nw21[(new BigNumber(16)).toNumber()] = _154_v72;
            _nw21[(new BigNumber(17)).toNumber()] = _154_v72;
            _nw21[(new BigNumber(18)).toNumber()] = (_154_v72).update(_150_v68, _153_v71);
            _nw21[(new BigNumber(19)).toNumber()] = ((_154_v72).update(_150_v68, _153_v71)).update(_150_v68, _153_v71);
            _nw21[(new BigNumber(20)).toNumber()] = (((_157_v75).contains(_module.__default.fm16(p2, globalState))) ? ((_157_v75).get(_module.__default.fm16(p2, globalState))) : (_154_v72));
            _158_v76 = _nw21;
            let _index19 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((_158_v76).length));
            (_158_v76)[_index19] = _154_v72;
            let _index20 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((_158_v76).length));
            (_158_v76)[_index20] = (_154_v72).Merge((_154_v72).Merge(_154_v72));
            let _159_v77;
            _159_v77 = _dafny.Map.Empty.slice().updateUnsafe(true,_151_v69);
            _159_v77 = (_159_v77).update(_module.__default.fm24(p2, new BigNumber((_dafny.Seq.update(_147_v66, _module.__default.safeIndex((_150_v68).f7, new BigNumber((_147_v66).length)), _144_v63)).length), _147_v66, globalState), _151_v69);
            let _nw22 = Array((new BigNumber(23)).toNumber()).fill(false);
            _145_v64 = _nw22;
            let _160_v78;
            _160_v78 = _dafny.Map.Empty.slice().updateUnsafe(_150_v68,_151_v69);
            _160_v78 = _dafny.Map.Empty.slice().updateUnsafe(_150_v68,_151_v69);
          }
        }
      }
      let _index21 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_145_v64).length));
      (_145_v64)[_index21] = _dafny.Seq.IsPrefixOf(_147_v66, _147_v66);
      let _index22 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_145_v64).length));
      (_145_v64)[_index22] = (p0).isLessThanOrEqualTo(p0);
      let _161_v79;
      let _nw23 = Array((new BigNumber(3)).toNumber()).fill([]);
      _161_v79 = _nw23;
      let _162_v80;
      let _nw24 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
      _162_v80 = _nw24;
      let _index23 = _module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length));
      (_161_v79)[_index23] = _162_v80;
      let _index24 = _module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length));
      (_161_v79)[_index24] = _162_v80;
      if ((_145_v64)[_module.__default.safeIndex(new BigNumber(630), new BigNumber((_145_v64).length))]) {
        let _163_v81;
        _163_v81 = _module.D7.create_DC19(p0, (_145_v64)[_module.__default.safeIndex(new BigNumber(630), new BigNumber((_145_v64).length))], p0);
        let _164_v82;
        _164_v82 = _dafny.Seq.of(_163_v81);
        let _165_v83;
        _165_v83 = _dafny.MultiSet.fromElements(p0, p0, new BigNumber((_164_v82).length), p0, new BigNumber((_147_v66).length));
        _165_v83 = _165_v83;
        let _arr0 = (_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))];
        let _index25 = _module.__default.safeIndex(new BigNumber(980), new BigNumber(((_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))]).length));
        _arr0[_index25] = new BigNumber(237);
        let _166_v84;
        _166_v84 = _dafny.MultiSet.fromElements(p2, (_145_v64)[_module.__default.safeIndex(new BigNumber(630), new BigNumber((_145_v64).length))]);
        let _167_v85;
        let _nw25 = new _module.C1();
        _nw25.__ctor(p0, p0);
        _167_v85 = _nw25;
        let _168_v86;
        _168_v86 = _dafny.Map.Empty.slice().updateUnsafe(_167_v85,p1);
        let _169_v87;
        _169_v87 = _167_v85;
        let _170_v88;
        _170_v88 = _module.D9.create_DC23((_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))]);
        let _171_v89;
        _171_v89 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_170_v88).dtor_cf36);
        let _172_v90;
        let _nw26 = Array((new BigNumber(18)).toNumber()).fill(_module.D9.Default());
        _172_v90 = _nw26;
        let _173_v91;
        _173_v91 = _dafny.Map.Empty.slice().updateUnsafe(p1,_172_v90);
        let _index26 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_145_v64).length));
        let _index27 = _module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length));
        let _arr1 = (_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))];
        let _index28 = _module.__default.safeIndex(new BigNumber(980), new BigNumber(((_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))]).length));
        let _rhs11 = (_166_v84).IsDisjointFrom(_166_v84);
        let _rhs12 = (((_168_v86).contains((_169_v87))) ? ((_168_v86).get((_169_v87))) : (true));
        let _rhs13 = (((_171_v89).contains(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(p2,_172_v90)).Merge(_173_v91)).length))) ? ((_171_v89).get(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(p2,_172_v90)).Merge(_173_v91)).length))) : (((p1) ? ((_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))]) : (_162_v80))));
        let _rhs14 = false;
        let _rhs15 = (p0).minus(new BigNumber(-541));
        let _lhs12 = _145_v64;
        let _lhs13 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_145_v64).length));
        let _lhs14 = globalState;
        let _lhs15 = _161_v79;
        let _lhs16 = _module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length));
        let _lhs17 = globalState;
        let _lhs18 = (_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))];
        let _lhs19 = _module.__default.safeIndex(new BigNumber(980), new BigNumber(((_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))]).length));
        _lhs12[_lhs13] = _rhs11;
        _lhs14.f3 = _rhs12;
        _lhs15[_lhs16] = _rhs13;
        _lhs17.f4 = _rhs14;
        _lhs18[_lhs19] = _rhs15;
        (globalState).f4 = true;
        let _174_v92;
        _174_v92 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.update(_56_v0, _module.__default.safeIndex(new BigNumber(-735), new BigNumber((_56_v0).length)), p0), _dafny.Seq.of(((_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))])[_module.__default.safeIndex(new BigNumber(980), new BigNumber(((_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))]).length))])),p0);
        _174_v92 = ((_174_v92).Merge(_174_v92)).Merge(_174_v92);
        let _175_v93;
        let _nw27 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.of());
        _175_v93 = _nw27;
        _175_v93 = ((true) ? (_175_v93) : (_175_v93));
      } else {
        let _176_v94;
        _176_v94 = new BigNumber(-106);
        let _177_v95;
        _177_v95 = _dafny.MultiSet.fromElements(p2);
        let _rhs16 = false;
        let _rhs17 = _176_v94;
        let _rhs18 = ((_dafny.ZERO).minus(_176_v94)).isEqualTo(p0);
        let _rhs19 = (p0).isLessThanOrEqualTo(_module.__default.safeModuloInt(new BigNumber(265), (((_177_v95).contains(p2)) ? ((_177_v95).get(p2)) : ((_dafny.ZERO).minus(_176_v94)))));
        let _lhs20 = globalState;
        let _lhs21 = globalState;
        let _lhs22 = globalState;
        _lhs20.f4 = _rhs16;
        _176_v94 = _rhs17;
        _lhs21.f4 = _rhs18;
        _lhs22.f3 = _rhs19;
        (globalState).f4 = p2;
        let _178_v96;
        _178_v96 = _dafny.Map.Empty.slice().updateUnsafe(!(new BigNumber((_147_v66).length)).isEqualTo(p0),_dafny.MultiSet.FromArray(_56_v0));
        _178_v96 = (_178_v96).update(p2, _dafny.MultiSet.FromArray(_56_v0));
        let _arr2 = (_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))];
        let _index29 = _module.__default.safeIndex(new BigNumber(869), new BigNumber(((_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))]).length));
        _arr2[_index29] = p0;
        let _arr3 = (_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))];
        let _index30 = _module.__default.safeIndex(new BigNumber(869), new BigNumber(((_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))]).length));
        _arr3[_index30] = (_module.__default.fm16(p1, globalState)).minus(p0);
        let _179_v97;
        _179_v97 = _module.D1.create_DC1(new _dafny.CodePoint('d'.codePointAt(0)));
        let _source5 = _module.D1.create_DC3(_179_v97);
        if (_source5.is_DC2) {
          let _nw28 = Array((new BigNumber(7)).toNumber());
          _nw28[(_dafny.ZERO).toNumber()] = _147_v66;
          _nw28[(_dafny.ONE).toNumber()] = _dafny.Seq.UnicodeFromString("mfn");
          _nw28[(new BigNumber(2)).toNumber()] = ((!(p1)) ? (_147_v66) : (_147_v66));
          _nw28[(new BigNumber(3)).toNumber()] = _dafny.Seq.UnicodeFromString("ilvyw");
          _nw28[(new BigNumber(4)).toNumber()] = _147_v66;
          _nw28[(new BigNumber(5)).toNumber()] = _147_v66;
          _nw28[(new BigNumber(6)).toNumber()] = _147_v66;
          r0 = _nw28;
          _176_v94 = _176_v94;
          let _180_v98;
          _180_v98 = _module.D3.create_DC7(p1);
          _180_v98 = _180_v98;
          let _181_v99;
          let _nw29 = Array((new BigNumber(9)).toNumber()).fill([]);
          _181_v99 = _nw29;
          let _arr4 = (_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))];
          let _index31 = _module.__default.safeIndex(new BigNumber(869), new BigNumber(((_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))]).length));
          let _rhs20 = _181_v99;
          let _rhs21 = ((_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))])[_module.__default.safeIndex(new BigNumber(869), new BigNumber(((_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))]).length))];
          let _lhs23 = (_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))];
          let _lhs24 = _module.__default.safeIndex(new BigNumber(869), new BigNumber(((_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))]).length));
          _181_v99 = _rhs20;
          _lhs23[_lhs24] = _rhs21;
        } else if (_source5.is_DC1) {
          let _182___mcc_h8 = (_source5).cf1;
          let _183_cf1 = _182___mcc_h8;
          (globalState).f3 = !_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("lq"), _183_cf1);
          let _184_v100;
          let _nw30 = new _module.C5();
          _nw30.__ctor(_147_v66, ((_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))])[_module.__default.safeIndex(new BigNumber(869), new BigNumber(((_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))]).length))], p0, new BigNumber(780));
          _184_v100 = _nw30;
          let _185_v101;
          _185_v101 = _dafny.Map.Empty.slice().updateUnsafe(_184_v100,p0);
          let _186_v102;
          _186_v102 = _dafny.Map.Empty.slice().updateUnsafe((_185_v101).Merge(_185_v101),p1);
          _186_v102 = (_186_v102).update(_185_v101, p2);
          let _index32 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_145_v64).length));
          let _rhs22 = (_145_v64)[_module.__default.safeIndex(new BigNumber(630), new BigNumber((_145_v64).length))];
          let _rhs23 = _56_v0;
          let _lhs25 = _145_v64;
          let _lhs26 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_145_v64).length));
          _lhs25[_lhs26] = _rhs22;
          _56_v0 = _rhs23;
          let _187_v103;
          let _nw31 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.of());
          _187_v103 = _nw31;
          let _188_v104;
          _188_v104 = _dafny.Seq.of(_145_v64, _145_v64, _145_v64);
          let _index33 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_187_v103).length));
          (_187_v103)[_index33] = _188_v104;
          let _index34 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_187_v103).length));
          (_187_v103)[_index34] = _dafny.Seq.Concat(_dafny.Seq.Concat(_188_v104, _188_v104), _188_v104);
        } else {
          let _189___mcc_h9 = (_source5).cf2;
          let _190_cf2 = _189___mcc_h9;
          let _191_v105;
          let _nw32 = new _module.C6();
          _nw32.__ctor(new BigNumber((_56_v0).length), p0);
          _191_v105 = _nw32;
          let _192_v106;
          _192_v106 = _dafny.Seq.of(_191_v105, _191_v105);
          let _193_v107;
          _193_v107 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_192_v106).length),_176_v94);
          let _194_v108;
          let _nw33 = new _module.C4();
          _nw33.__ctor(p2, _176_v94, _176_v94);
          _194_v108 = _nw33;
          let _195_v109;
          _195_v109 = _dafny.Map.Empty.slice().updateUnsafe(_194_v108,((_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))])[_module.__default.safeIndex(new BigNumber(869), new BigNumber(((_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))]).length))]);
          let _196_v110;
          let _nw34 = new _module.C4();
          _nw34.__ctor(p1, p0, (((_195_v109).contains(_194_v108)) ? ((_195_v109).get(_194_v108)) : (new BigNumber(325))));
          _196_v110 = _nw34;
          _176_v94 = (_module.__default.safeDivisionInt((((_193_v107).contains(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_145_v64)[_module.__default.safeIndex(new BigNumber(630), new BigNumber((_145_v64).length))],_196_v110)).length))) ? ((_193_v107).get(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_145_v64)[_module.__default.safeIndex(new BigNumber(630), new BigNumber((_145_v64).length))],_196_v110)).length))) : (new BigNumber((_56_v0).length))), _176_v94)).plus((((_193_v107).contains(new BigNumber(64))) ? ((_193_v107).get(new BigNumber(64))) : (new BigNumber(-306))));
          let _197_v111;
          let _nw35 = new _module.C2();
          _nw35.__ctor(((_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))])[_module.__default.safeIndex(new BigNumber(869), new BigNumber(((_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))]).length))], ((_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))])[_module.__default.safeIndex(new BigNumber(869), new BigNumber(((_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))]).length))]);
          _197_v111 = _nw35;
          let _arr5 = (_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))];
          let _index35 = _module.__default.safeIndex(new BigNumber(869), new BigNumber(((_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))]).length));
          let _rhs24 = _145_v64;
          let _rhs25 = _dafny.areEqual(_dafny.Seq.Concat(_147_v66, _dafny.Seq.Create(_module.__default.abs(new BigNumber(616)), ((_198_v63) => function (_199_i4) {
            return _198_v63;
          })(_144_v63))), _147_v66);
          let _rhs26 = _176_v94;
          let _lhs27 = globalState;
          let _lhs28 = (_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))];
          let _lhs29 = _module.__default.safeIndex(new BigNumber(869), new BigNumber(((_161_v79)[_module.__default.safeIndex(new BigNumber(944), new BigNumber((_161_v79).length))]).length));
          _145_v64 = _rhs24;
          _lhs27.f3 = _rhs25;
          _lhs28[_lhs29] = _rhs26;
          let _200_v112;
          let _nw36 = Array((new BigNumber(19)).toNumber());
          _200_v112 = _nw36;
          let _201_v113;
          let _nw37 = new _module.C3();
          _nw37.__ctor(p1, p0, p0);
          _201_v113 = _nw37;
          let _index36 = _module.__default.safeIndex(new BigNumber(994), new BigNumber((_200_v112).length));
          (_200_v112)[_index36] = _201_v113;
          let _index37 = _module.__default.safeIndex(new BigNumber(994), new BigNumber((_200_v112).length));
          (_200_v112)[_index37] = _201_v113;
        }
      }
      let _202_v114;
      _202_v114 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_147_v66),!((_145_v64)[_module.__default.safeIndex(new BigNumber(630), new BigNumber((_145_v64).length))]));
      let _203_v115;
      _203_v115 = _dafny.MultiSet.fromElements(_147_v66);
      let _204_v116;
      let _nw38 = new _module.C0();
      _nw38.__ctor(_162_v80, p0, p0);
      _204_v116 = _nw38;
      let _205_v117;
      _205_v117 = _dafny.Map.Empty.slice().updateUnsafe(p0,_204_v116);
      let _206_v118;
      _206_v118 = _dafny.Map.Empty.slice().updateUnsafe(p0,_205_v117);
      _202_v114 = (_202_v114).update(_203_v115, (_206_v118).contains(p0));
      let _207_v119;
      let _nw39 = Array((new BigNumber(9)).toNumber());
      _nw39[(_dafny.ZERO).toNumber()] = _147_v66;
      _nw39[(_dafny.ONE).toNumber()] = _147_v66;
      _nw39[(new BigNumber(2)).toNumber()] = _147_v66;
      _nw39[(new BigNumber(3)).toNumber()] = _147_v66;
      _nw39[(new BigNumber(4)).toNumber()] = _147_v66;
      _nw39[(new BigNumber(5)).toNumber()] = _147_v66;
      _nw39[(new BigNumber(6)).toNumber()] = _147_v66;
      _nw39[(new BigNumber(7)).toNumber()] = _147_v66;
      _nw39[(new BigNumber(8)).toNumber()] = _147_v66;
      _207_v119 = _nw39;
      let _208_v120;
      _208_v120 = _module.D14.create_DC35(_207_v119);
      r0 = (_208_v120).dtor_cf48;
      let _209_v121;
      _209_v121 = _module.D5.create_DC13(p2, _144_v63);
      let _pat_let_tv0 = _147_v66;
      let _pat_let_tv1 = p2;
      r1 = function (_source6) {
        if (_source6.is_DC12) {
          let _210___mcc_h10 = (_source6).cf19;
          let _211_cf19 = _210___mcc_h10;
          return _dafny.Seq.IsPrefixOf(_pat_let_tv0, (_module.D7.create_DC17(_dafny.Seq.UnicodeFromString("v"))).dtor_cf26);
        } else if (_source6.is_DC13) {
          let _212___mcc_h11 = (_source6).cf20;
          let _213___mcc_h12 = (_source6).cf21;
          let _214_cf21 = _213___mcc_h12;
          let _215_cf20 = _212___mcc_h11;
          return _215_cf20;
        } else if (_source6.is_DC11) {
          let _216___mcc_h13 = (_source6).cf18;
          let _217_cf18 = _216___mcc_h13;
          return _pat_let_tv1;
        } else {
          let _218___mcc_h14 = (_source6).cf22;
          let _219_cf22 = _218___mcc_h14;
          return true;
        }
      }(_209_v121);
      return [r0, r1];
    }
    static Main(__noArgsParameter) {
      let _220_v0;
      let _init1 = function (_221_i1) {
        return false;
      };
      let _nw40 = Array((new BigNumber(18)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw40.length); _i0_1++) {
        _nw40[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _220_v0 = _nw40;
      let _222_v1;
      _222_v1 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(230)), function (_223_i0) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      }),_220_v0);
      let _224_v2;
      _224_v2 = _dafny.Seq.UnicodeFromString("w");
      let _225_globalState;
      let _nw41 = new _module.GlobalState();
      _nw41.__ctor(new BigNumber(124), true, (_222_v1).Merge(_222_v1), false, true, new BigNumber(611), _224_v2);
      _225_globalState = _nw41;
      let _226_v3;
      _226_v3 = false;
      (_225_globalState).f3 = _226_v3;
      let _227_v4;
      _227_v4 = new BigNumber(756);
      let _228_i2;
      _228_i2 = _dafny.ZERO;
      L1: {
        while ((_module.__default.safeModuloInt(_227_v4, _227_v4)).isLessThanOrEqualTo(_227_v4)) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_228_i2)) {
              break L1;
            }
            _228_i2 = (_228_i2).plus(_dafny.ONE);
            let _229_v5;
            let _230_v6;
            let _out0;
            let _out1;
            let _outcollector0 = _module.__default.m0(_227_v4, _226_v3, _226_v3, _225_globalState);
            _out0 = _outcollector0[0];
            _out1 = _outcollector0[1];
            _229_v5 = _out0;
            _230_v6 = _out1;
            let _231_v7;
            _231_v7 = _dafny.Map.Empty.slice().updateUnsafe(true,_230_v6);
            let _232_v8;
            _232_v8 = _dafny.Seq.of(_226_v3);
            _231_v7 = (_231_v7).update(_dafny.Seq.contains(_232_v8, _226_v3), _226_v3);
            let _233_v9;
            _233_v9 = _dafny.Set.fromElements(_224_v2);
            let _234_v10;
            let _nw42 = new _module.C2();
            _nw42.__ctor((new BigNumber((_233_v9).length)).minus(_227_v4), (_dafny.ZERO).minus(_227_v4));
            _234_v10 = _nw42;
            _231_v7 = (_231_v7).update(_226_v3, _226_v3);
          }
        }
      }
      let _235_v11;
      _235_v11 = _dafny.Seq.of(new BigNumber(88));
      (_225_globalState).f4 = ((_235_v11)[_module.__default.safeIndex((_dafny.ZERO).minus(_227_v4), new BigNumber((_235_v11).length))]).isEqualTo(new BigNumber((_224_v2).length));
      let _236_v12;
      _236_v12 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(823),false);
      let _237_v13;
      let _nw43 = new _module.C4();
      _nw43.__ctor(_226_v3, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_226_v3,_226_v3)).length), _227_v4);
      _237_v13 = _nw43;
      let _238_v14;
      _238_v14 = _dafny.Seq.of(_237_v13, _237_v13);
      let _239_v15;
      let _nw44 = Array((new BigNumber(23)).toNumber());
      _nw44[(_dafny.ZERO).toNumber()] = new BigNumber(218);
      _nw44[(_dafny.ONE).toNumber()] = new BigNumber((_224_v2).length);
      _nw44[(new BigNumber(2)).toNumber()] = _227_v4;
      _nw44[(new BigNumber(3)).toNumber()] = new BigNumber((_236_v12).length);
      _nw44[(new BigNumber(4)).toNumber()] = _227_v4;
      _nw44[(new BigNumber(5)).toNumber()] = new BigNumber(709);
      _nw44[(new BigNumber(6)).toNumber()] = _227_v4;
      _nw44[(new BigNumber(7)).toNumber()] = new BigNumber((_238_v14).length);
      _nw44[(new BigNumber(8)).toNumber()] = _227_v4;
      _nw44[(new BigNumber(9)).toNumber()] = _227_v4;
      _nw44[(new BigNumber(10)).toNumber()] = _227_v4;
      _nw44[(new BigNumber(11)).toNumber()] = new BigNumber((_dafny.Seq.of((_237_v13).f11, (_237_v13).f11)).length);
      _nw44[(new BigNumber(12)).toNumber()] = _227_v4;
      _nw44[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus(_227_v4);
      _nw44[(new BigNumber(14)).toNumber()] = _dafny.ZERO;
      _nw44[(new BigNumber(15)).toNumber()] = _227_v4;
      _nw44[(new BigNumber(16)).toNumber()] = _227_v4;
      _nw44[(new BigNumber(17)).toNumber()] = _227_v4;
      _nw44[(new BigNumber(18)).toNumber()] = _227_v4;
      _nw44[(new BigNumber(19)).toNumber()] = _227_v4;
      _nw44[(new BigNumber(20)).toNumber()] = _227_v4;
      _nw44[(new BigNumber(21)).toNumber()] = _227_v4;
      _nw44[(new BigNumber(22)).toNumber()] = _227_v4;
      _239_v15 = _nw44;
      let _240_v16;
      _240_v16 = _dafny.Seq.of(_239_v15);
      let _241_v17;
      let _nw45 = new _module.C0();
      _nw45.__ctor((_240_v16)[_module.__default.safeIndex(_227_v4, new BigNumber((_240_v16).length))], new BigNumber(525), _227_v4);
      _241_v17 = _nw45;
      let _index38 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_220_v0).length));
      (_220_v0)[_index38] = !(!(!((_237_v13).f11) || ((_237_v13).f11)));
      let _242_v18;
      let _nw46 = Array((_dafny.ONE).toNumber());
      _nw46[(_dafny.ZERO).toNumber()] = _239_v15;
      _242_v18 = _nw46;
      let _index39 = _module.__default.safeIndex(new BigNumber(677), new BigNumber((_242_v18).length));
      (_242_v18)[_index39] = _241_v17.f13;
      let _243_v19;
      _243_v19 = _dafny.Seq.of((_237_v13).f11, _dafny.areEqual(_module.__default.fm8(true, _235_v11, _226_v3, _225_globalState), _224_v2), ((!((_237_v13).f11)) ? (_226_v3) : ((_237_v13).f11)), !(_236_v12).equals(_236_v12), ((_237_v13).f11) === (_226_v3));
      let _244_v20;
      _244_v20 = new _dafny.CodePoint('b'.codePointAt(0));
      let _245_v21;
      _245_v21 = _module.D1.create_DC1(_244_v20);
      let _246_v22;
      _246_v22 = _module.D1.create_DC3(_module.D1.create_DC1(_244_v20));
      let _247_v23;
      _247_v23 = _dafny.MultiSet.fromElements(_module.D1.create_DC3(_245_v21), _246_v22);
      let _248_v24;
      _248_v24 = _dafny.Set.fromElements(_247_v23, _247_v23, _dafny.MultiSet.fromElements(_246_v22));
      let _index40 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_220_v0).length));
      let _index41 = _module.__default.safeIndex(new BigNumber(677), new BigNumber((_242_v18).length));
      let _rhs27 = !((((new BigNumber((_224_v2).length)).isLessThanOrEqualTo(_227_v4)) ? (_226_v3) : (_226_v3)));
      let _rhs28 = _239_v15;
      let _rhs29 = _243_v19;
      let _rhs30 = _module.__default.fm20(_226_v3, (_248_v24).IsProperSubsetOf(_module.__default.fm36(_226_v3, (_237_v13).f11, _225_globalState)), (_237_v13).f11, _225_globalState);
      let _rhs31 = _237_v13;
      let _lhs30 = _220_v0;
      let _lhs31 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_220_v0).length));
      let _lhs32 = _242_v18;
      let _lhs33 = _module.__default.safeIndex(new BigNumber(677), new BigNumber((_242_v18).length));
      _lhs30[_lhs31] = _rhs27;
      _lhs32[_lhs33] = _rhs28;
      _243_v19 = _rhs29;
      _227_v4 = _rhs30;
      _237_v13 = _rhs31;
      let _249_v25;
      _249_v25 = _dafny.MultiSet.fromElements(_241_v17.f13);
      let _250_v26;
      _250_v26 = _dafny.Set.fromElements((_237_v13).f11);
      let _251_v27;
      _251_v27 = _dafny.Map.Empty.slice().updateUnsafe((_237_v13).f11,(_241_v17).fm1(_250_v26, _227_v4, _227_v4, _227_v4, _225_globalState));
      let _252_v28;
      _252_v28 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_227_v4),(_249_v25).update(_241_v17.f13, _module.__default.abs(_227_v4)));
      let _rhs32 = _220_v0;
      let _rhs33 = (((((_251_v27).contains((_237_v13).f11)) ? ((_251_v27).get((_237_v13).f11)) : ((_237_v13).f11))) ? (new BigNumber(((_250_v26).Difference(_250_v26)).length)) : (_227_v4));
      let _rhs34 = (((_252_v28).contains((_dafny.ZERO).minus(_227_v4))) ? ((_252_v28).get((_dafny.ZERO).minus(_227_v4))) : (_249_v25));
      _220_v0 = _rhs32;
      _227_v4 = _rhs33;
      _249_v25 = _rhs34;
      let _253_v29;
      let _nw47 = new _module.C2();
      _nw47.__ctor(_227_v4, _227_v4);
      _253_v29 = _nw47;
      let _254_v30;
      _254_v30 = _dafny.Map.Empty.slice().updateUnsafe(_253_v29,_236_v12);
      let _255_v31;
      let _nw48 = new _module.C2();
      _nw48.__ctor(_227_v4, (new BigNumber(356)).plus(new BigNumber(((((_254_v30).contains(_253_v29)) ? ((_254_v30).get(_253_v29)) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_253_v29).f8,_241_v17)).length),(_220_v0)[_module.__default.safeIndex(new BigNumber(913), new BigNumber((_220_v0).length))])))).length)));
      _255_v31 = _nw48;
      let _256_i3;
      _256_i3 = _dafny.ZERO;
      L2: {
        while (_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_224_v2, _224_v2), _dafny.Seq.Create(_module.__default.abs(new BigNumber(683)), ((_265_v20) => function (_266_i4) {
          return _265_v20;
        })(_244_v20)))) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_256_i3)) {
              break L2;
            }
            _256_i3 = (_256_i3).plus(_dafny.ONE);
            let _index42 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_220_v0).length));
            (_220_v0)[_index42] = !((_237_v13).f11);
            let _257_v32;
            let _nw49 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.of());
            _257_v32 = _nw49;
            let _258_v33;
            _258_v33 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements((_220_v0)[_module.__default.safeIndex(new BigNumber(913), new BigNumber((_220_v0).length))]),(_253_v29).f7);
            let _259_v34;
            _259_v34 = _module.D6.create_DC16(_258_v33, _243_v19);
            let _index43 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_257_v32).length));
            (_257_v32)[_index43] = (_259_v34).dtor_cf25;
            let _index44 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_257_v32).length));
            (_257_v32)[_index44] = _243_v19;
            let _260_v35;
            _260_v35 = _module.D7.create_DC20(_244_v20, _227_v4);
            let _261_v36;
            _261_v36 = _dafny.Map.Empty.slice().updateUnsafe(_260_v35,_244_v20);
            let _262_v37;
            let _nw50 = new _module.C4();
            _nw50.__ctor(_226_v3, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_261_v36, _261_v36)).length)), (_253_v29).f7);
            _262_v37 = _nw50;
            _224_v2 = (((_220_v0)[_module.__default.safeIndex(new BigNumber(913), new BigNumber((_220_v0).length))]) ? (_dafny.Seq.Concat((_255_v31).fm14(_225_globalState), _224_v2)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(220)), ((_263_v20) => function (_264_i5) {
              return _263_v20;
            })(_244_v20))));
          }
        }
      }
      let _267_v38;
      let _init2 = ((_268_v0, _269_v13) => function (_270_i6) {
        return _dafny.Seq.of(_dafny.Set.fromElements((_268_v0)[_module.__default.safeIndex(new BigNumber(913), new BigNumber((_268_v0).length))], (_269_v13).f11));
      })(_220_v0, _237_v13);
      let _nw51 = Array((new BigNumber(5)).toNumber());
      for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw51.length); _i0_2++) {
        _nw51[_i0_2] = _init2(new BigNumber(_i0_2));
      }
      _267_v38 = _nw51;
      let _271_v39;
      _271_v39 = _dafny.Seq.of(_250_v26);
      let _index45 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_267_v38).length));
      (_267_v38)[_index45] = _271_v39;
      let _272_v40;
      _272_v40 = _module.D4.create_DC10(_227_v4, (_237_v13).f11);
      let _index46 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_220_v0).length));
      let _index47 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_267_v38).length));
      let _rhs35 = (_272_v40).dtor_cf17;
      let _rhs36 = _dafny.Seq.Concat(_271_v39, _271_v39);
      let _lhs34 = _220_v0;
      let _lhs35 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_220_v0).length));
      let _lhs36 = _267_v38;
      let _lhs37 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_267_v38).length));
      _lhs34[_lhs35] = _rhs35;
      _lhs36[_lhs37] = _rhs36;
      let _273_v41;
      let _nw52 = new _module.C5();
      _nw52.__ctor(_224_v2, _227_v4, (_253_v29).f8, (_253_v29).f7);
      _273_v41 = _nw52;
      let _274_v42;
      _274_v42 = _dafny.Map.Empty.slice().updateUnsafe(_273_v41,(_273_v41).f10);
      let _275_v43;
      _275_v43 = _dafny.MultiSet.fromElements(_220_v0, _220_v0);
      (_225_globalState).f4 = (_237_v13).fm1(_250_v26, (_253_v29).f8, (((_274_v42).contains(_273_v41)) ? ((_274_v42).get(_273_v41)) : (_227_v4)), (((_275_v43).contains(_220_v0)) ? ((_275_v43).get(_220_v0)) : (new BigNumber((_235_v11).length))), _225_globalState);
      (_225_globalState).f3 = !((((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update((_273_v41).f9, _module.__default.safeIndex(new BigNumber((_250_v26).length), new BigNumber(((_273_v41).f9).length)), _244_v20)).length))).multipliedBy((_253_v29).f7)).isLessThan(new BigNumber((_module.__default.fm9(_226_v3, _225_globalState)).length)));
      let _hi0 = _227_v4;
      for (let _276_i7 = ((_253_v29).f7).minus(new BigNumber((_235_v11).length)); _276_i7.isLessThan(_hi0); _276_i7 = _276_i7.plus(_dafny.ONE)) {
        _227_v4 = _227_v4;
        let _277_v44;
        _277_v44 = _dafny.Seq.of(_253_v29, _253_v29);
        let _278_v45;
        let _nw53 = new _module.C2();
        _nw53.__ctor((_253_v29).f7, new BigNumber((_dafny.Seq.Concat(_277_v44, _277_v44)).length));
        _278_v45 = _nw53;
        _278_v45 = _278_v45;
        _278_v45 = _278_v45;
        (_278_v45).m1(_224_v2, _225_globalState);
      }
      let _279_v46;
      let _nw54 = Array((new BigNumber(9)).toNumber()).fill(_module.D8.Default());
      _279_v46 = _nw54;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_279_v46).length))) {
        let _280_i8 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_280_i8)) && ((_280_i8).isLessThan(new BigNumber((_279_v46).length))))) {
          (_279_v46)[(_280_i8)] = _module.D8.create_DC22((_dafny.ZERO).minus((_273_v41).f10));
        }
      }
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_241_v17.f13).length))) {
        let _281_i9 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_281_i9)) && ((_281_i9).isLessThan(new BigNumber((_241_v17.f13).length))))) {
          let _arr6 = _241_v17.f13;
          _arr6[(_281_i9)] = _module.__default.safeDivisionInt(_281_i9, (_253_v29).f7);
        }
      }
      (_225_globalState).f3 = !(_226_v3);
      _224_v2 = _224_v2;
      process.stdout.write(_dafny.toString((_220_v0)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v0)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v0)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v0)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v0)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v0)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v0)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v0)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v0)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v0)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v0)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v0)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v0)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v0)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v0)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v0)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v0)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v0)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_222_v1).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_224_v2).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_225_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_225_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_225_globalState).f2).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_225_globalState.f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_225_globalState.f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_225_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_225_globalState).f6).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_226_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_227_v4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_228_i2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_235_v11, _dafny.Seq.of(new BigNumber(88)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_236_v12).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(823),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_237_v13).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_238_v14).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v15)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v15)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v15)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v15)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v15)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v15)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v15)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v15)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v15)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v15)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v15)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v15)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v15)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v15)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v15)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v15)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v15)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v15)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v15)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v15)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v15)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v15)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v15)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_240_v16).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_241_v17.f13)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_241_v17.f13)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_241_v17.f13)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_241_v17.f13)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_241_v17.f13)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_241_v17.f13)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_241_v17.f13)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_241_v17.f13)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_241_v17.f13)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_241_v17.f13)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_241_v17.f13)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_241_v17.f13)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_241_v17.f13)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_241_v17.f13)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_241_v17.f13)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_241_v17.f13)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_241_v17.f13)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_241_v17.f13)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_241_v17.f13)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_241_v17.f13)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_241_v17.f13)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_241_v17.f13)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_241_v17.f13)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_242_v18)[_dafny.ZERO])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_242_v18)[_dafny.ZERO])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_242_v18)[_dafny.ZERO])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_242_v18)[_dafny.ZERO])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_242_v18)[_dafny.ZERO])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_242_v18)[_dafny.ZERO])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_242_v18)[_dafny.ZERO])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_242_v18)[_dafny.ZERO])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_242_v18)[_dafny.ZERO])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_242_v18)[_dafny.ZERO])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_242_v18)[_dafny.ZERO])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_242_v18)[_dafny.ZERO])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_242_v18)[_dafny.ZERO])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_242_v18)[_dafny.ZERO])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_242_v18)[_dafny.ZERO])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_242_v18)[_dafny.ZERO])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_242_v18)[_dafny.ZERO])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_242_v18)[_dafny.ZERO])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_242_v18)[_dafny.ZERO])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_242_v18)[_dafny.ZERO])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_242_v18)[_dafny.ZERO])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_242_v18)[_dafny.ZERO])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_242_v18)[_dafny.ZERO])[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_243_v19, _dafny.Seq.of(false, false, false, false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_244_v20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_245_v21).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_246_v22).dtor_cf2).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_247_v23).equals(_dafny.MultiSet.fromElements(_module.D1.create_DC3(_module.D1.create_DC1(new _dafny.CodePoint('b'.codePointAt(0)))), _module.D1.create_DC3(_module.D1.create_DC1(new _dafny.CodePoint('b'.codePointAt(0))))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_248_v24).equals(_dafny.Set.fromElements(_dafny.MultiSet.fromElements(_module.D1.create_DC3(_module.D1.create_DC1(new _dafny.CodePoint('b'.codePointAt(0)))), _module.D1.create_DC3(_module.D1.create_DC1(new _dafny.CodePoint('b'.codePointAt(0))))), _dafny.MultiSet.fromElements(_module.D1.create_DC3(_module.D1.create_DC1(new _dafny.CodePoint('b'.codePointAt(0)))))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_249_v25).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_250_v26).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_251_v27).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_252_v28).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_253_v29).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_253_v29).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_254_v30).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_256_i3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_267_v38)[_dafny.ZERO], _dafny.Seq.of(_dafny.Set.fromElements(true, false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_267_v38)[_dafny.ONE], _dafny.Seq.of(_dafny.Set.fromElements(true, false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_267_v38)[new BigNumber(2)], _dafny.Seq.of(_dafny.Set.fromElements(true, false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_267_v38)[new BigNumber(3)], _dafny.Seq.of(_dafny.Set.fromElements(false), _dafny.Set.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_267_v38)[new BigNumber(4)], _dafny.Seq.of(_dafny.Set.fromElements(true, false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_271_v39, _dafny.Seq.of(_dafny.Set.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_272_v40).dtor_cf16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_272_v40).dtor_cf17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_273_v41).f9).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_273_v41).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_274_v42).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_275_v43).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_279_v46)[_dafny.ZERO]).dtor_cf35));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_279_v46)[_dafny.ONE]).dtor_cf35));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_279_v46)[new BigNumber(2)]).dtor_cf35));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_279_v46)[new BigNumber(3)]).dtor_cf35));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_279_v46)[new BigNumber(4)]).dtor_cf35));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_279_v46)[new BigNumber(5)]).dtor_cf35));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_279_v46)[new BigNumber(6)]).dtor_cf35));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_279_v46)[new BigNumber(7)]).dtor_cf35));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_279_v46)[new BigNumber(8)]).dtor_cf35));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC0(cf0) {
      let $dt = new D0(0);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC0() { return this.$tag === 0; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf0 === other.cf0;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return false;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D0.Default();
        }
      };
    }
  }

  $module.D1 = class D1 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC2() {
      let $dt = new D1(0);
      return $dt;
    }
    static create_DC1(cf1) {
      let $dt = new D1(1);
      $dt.cf1 = cf1;
      return $dt;
    }
    static create_DC3(cf2) {
      let $dt = new D1(2);
      $dt.cf2 = cf2;
      return $dt;
    }
    get is_DC2() { return this.$tag === 0; }
    get is_DC1() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC2";
      } else if (this.$tag === 1) {
        return "D1.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf2) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf1, other.cf1);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf2, other.cf2);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC2();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D1.Default();
        }
      };
    }
  }

  $module.D2 = class D2 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC5(cf4, cf5, cf6, cf7, cf8) {
      let $dt = new D2(0);
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC4(cf3) {
      let $dt = new D2(1);
      $dt.cf3 = cf3;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC4() { return this.$tag === 1; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf3() { return this.cf3; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC4" + "(" + _dafny.toString(this.cf3) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf4 === other.cf4 && _dafny.areEqual(this.cf5, other.cf5) && _dafny.areEqual(this.cf6, other.cf6) && _dafny.areEqual(this.cf7, other.cf7) && this.cf8 === other.cf8;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf3, other.cf3);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC5(false, _dafny.ZERO, _dafny.ZERO, _dafny.ZERO, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC7(cf10) {
      let $dt = new D3(0);
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC6(cf9) {
      let $dt = new D3(1);
      $dt.cf9 = cf9;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf9() { return this.cf9; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC7" + "(" + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC6" + "(" + _dafny.toString(this.cf9) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf10 === other.cf10;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf9, other.cf9);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC7(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D3.Default();
        }
      };
    }
  }

  $module.D4 = class D4 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC9(cf12, cf13, cf14, cf15) {
      let $dt = new D4(0);
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC10(cf16, cf17) {
      let $dt = new D4(1);
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC8(cf11) {
      let $dt = new D4(2);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC9" + "(" + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC8" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf12, other.cf12) && _dafny.areEqual(this.cf13, other.cf13) && _dafny.areEqual(this.cf14, other.cf14) && _dafny.areEqual(this.cf15, other.cf15);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf16, other.cf16) && this.cf17 === other.cf17;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf11 === other.cf11;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC9(_dafny.ZERO, _dafny.Seq.of(), _dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D4.Default();
        }
      };
    }
  }

  $module.D5 = class D5 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC12(cf19) {
      let $dt = new D5(0);
      $dt.cf19 = cf19;
      return $dt;
    }
    static create_DC13(cf20, cf21) {
      let $dt = new D5(1);
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC11(cf18) {
      let $dt = new D5(2);
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC14(cf22) {
      let $dt = new D5(3);
      $dt.cf22 = cf22;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get is_DC11() { return this.$tag === 2; }
    get is_DC14() { return this.$tag === 3; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf22() { return this.cf22; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC12" + "(" + _dafny.toString(this.cf19) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC11" + "(" + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 3) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf22) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf19 === other.cf19;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf20 === other.cf20 && _dafny.areEqual(this.cf21, other.cf21);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf18, other.cf18);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf22, other.cf22);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC12(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC16(cf24, cf25) {
      let $dt = new D6(0);
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC15(cf23) {
      let $dt = new D6(1);
      $dt.cf23 = cf23;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf23() { return this.cf23; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC15" + "(" + _dafny.toString(this.cf23) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf24, other.cf24) && _dafny.areEqual(this.cf25, other.cf25);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf23 === other.cf23;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC16(_dafny.Map.Empty, _dafny.Seq.of());
    }
    static Rtd() {
      return class {
        static get Default() {
          return D6.Default();
        }
      };
    }
  }

  $module.D7 = class D7 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC18(cf27, cf28) {
      let $dt = new D7(0);
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC19(cf29, cf30, cf31) {
      let $dt = new D7(1);
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      return $dt;
    }
    static create_DC20(cf32, cf33) {
      let $dt = new D7(2);
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC17(cf26) {
      let $dt = new D7(3);
      $dt.cf26 = cf26;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get is_DC19() { return this.$tag === 1; }
    get is_DC20() { return this.$tag === 2; }
    get is_DC17() { return this.$tag === 3; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf26() { return this.cf26; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC18" + "(" + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC19" + "(" + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC20" + "(" + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 3) {
        return "D7.DC17" + "(" + this.cf26.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf27 === other.cf27 && this.cf28 === other.cf28;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf29, other.cf29) && this.cf30 === other.cf30 && _dafny.areEqual(this.cf31, other.cf31);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf32, other.cf32) && _dafny.areEqual(this.cf33, other.cf33);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf26, other.cf26);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC18(false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D7.Default();
        }
      };
    }
  }

  $module.D8 = class D8 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC22(cf35) {
      let $dt = new D8(0);
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC21(cf34) {
      let $dt = new D8(1);
      $dt.cf34 = cf34;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get is_DC21() { return this.$tag === 1; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf34() { return this.cf34; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC22" + "(" + _dafny.toString(this.cf35) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC21" + "(" + _dafny.toString(this.cf34) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf35, other.cf35);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf34 === other.cf34;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC22(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D8.Default();
        }
      };
    }
  }

  $module.D9 = class D9 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC24(cf37, cf38) {
      let $dt = new D9(0);
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC23(cf36) {
      let $dt = new D9(1);
      $dt.cf36 = cf36;
      return $dt;
    }
    get is_DC24() { return this.$tag === 0; }
    get is_DC23() { return this.$tag === 1; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf36() { return this.cf36; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC24" + "(" + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC23" + "(" + _dafny.toString(this.cf36) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf37, other.cf37) && _dafny.areEqual(this.cf38, other.cf38);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf36 === other.cf36;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC24(_dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D9.Default();
        }
      };
    }
  }

  $module.D10 = class D10 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC26(cf40) {
      let $dt = new D10(0);
      $dt.cf40 = cf40;
      return $dt;
    }
    static create_DC25(cf39) {
      let $dt = new D10(1);
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC27(cf41) {
      let $dt = new D10(2);
      $dt.cf41 = cf41;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get is_DC25() { return this.$tag === 1; }
    get is_DC27() { return this.$tag === 2; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf41() { return this.cf41; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC26" + "(" + _dafny.toString(this.cf40) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC25" + "(" + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 2) {
        return "D10.DC27" + "(" + _dafny.toString(this.cf41) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf40, other.cf40);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf39, other.cf39);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf41, other.cf41);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC26(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D10.Default();
        }
      };
    }
  }

  $module.D11 = class D11 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC29(cf43) {
      let $dt = new D11(0);
      $dt.cf43 = cf43;
      return $dt;
    }
    static create_DC30() {
      let $dt = new D11(1);
      return $dt;
    }
    static create_DC28(cf42) {
      let $dt = new D11(2);
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC31(cf44) {
      let $dt = new D11(3);
      $dt.cf44 = cf44;
      return $dt;
    }
    get is_DC29() { return this.$tag === 0; }
    get is_DC30() { return this.$tag === 1; }
    get is_DC28() { return this.$tag === 2; }
    get is_DC31() { return this.$tag === 3; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf44() { return this.cf44; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC29" + "(" + _dafny.toString(this.cf43) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC30";
      } else if (this.$tag === 2) {
        return "D11.DC28" + "(" + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 3) {
        return "D11.DC31" + "(" + _dafny.toString(this.cf44) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf43, other.cf43);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf42 === other.cf42;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf44, other.cf44);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC29(_dafny.Map.Empty);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D11.Default();
        }
      };
    }
  }

  $module.D12 = class D12 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC33(cf46) {
      let $dt = new D12(0);
      $dt.cf46 = cf46;
      return $dt;
    }
    static create_DC32(cf45) {
      let $dt = new D12(1);
      $dt.cf45 = cf45;
      return $dt;
    }
    get is_DC33() { return this.$tag === 0; }
    get is_DC32() { return this.$tag === 1; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf45() { return this.cf45; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC33" + "(" + this.cf46.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC32" + "(" + _dafny.toString(this.cf45) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf46, other.cf46);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf45, other.cf45);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC33(_dafny.Seq.UnicodeFromString(""));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D12.Default();
        }
      };
    }
  }

  $module.D13 = class D13 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC34(cf47) {
      let $dt = new D13(0);
      $dt.cf47 = cf47;
      return $dt;
    }
    get is_DC34() { return this.$tag === 0; }
    get dtor_cf47() { return this.cf47; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC34" + "(" + _dafny.toString(this.cf47) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf47 === other.cf47;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return null;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D13.Default();
        }
      };
    }
  }

  $module.D14 = class D14 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC36(cf49, cf50, cf51, cf52, cf53) {
      let $dt = new D14(0);
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      return $dt;
    }
    static create_DC35(cf48) {
      let $dt = new D14(1);
      $dt.cf48 = cf48;
      return $dt;
    }
    static create_DC37(cf54) {
      let $dt = new D14(2);
      $dt.cf54 = cf54;
      return $dt;
    }
    get is_DC36() { return this.$tag === 0; }
    get is_DC35() { return this.$tag === 1; }
    get is_DC37() { return this.$tag === 2; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf54() { return this.cf54; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC36" + "(" + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ", " + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ", " + _dafny.toString(this.cf53) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC35" + "(" + _dafny.toString(this.cf48) + ")";
      } else if (this.$tag === 2) {
        return "D14.DC37" + "(" + _dafny.toString(this.cf54) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf49 === other.cf49 && _dafny.areEqual(this.cf50, other.cf50) && _dafny.areEqual(this.cf51, other.cf51) && _dafny.areEqual(this.cf52, other.cf52) && _dafny.areEqual(this.cf53, other.cf53);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf48 === other.cf48;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf54, other.cf54);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC36(false, _dafny.Seq.of(), _dafny.Seq.of(), _dafny.Seq.of(), new _dafny.CodePoint('D'.codePointAt(0)));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D14.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.T1 = class T1 {
  };

  $module.T2 = class T2 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f3 = false;
      this.f4 = false;
      this._f0 = _dafny.ZERO;
      this._f1 = false;
      this._f2 = _dafny.Map.Empty;
      this._f5 = _dafny.ZERO;
      this._f6 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this).f3 = f3;
      (_this).f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      return;
    }
    get f0() {
      let _this = this;
      return _this._f0;
    };
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f2() {
      let _this = this;
      return _this._f2;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f7 = _dafny.ZERO;
      this._f8 = _dafny.ZERO;
      this.f13 = [];
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    __ctor(f13, f7, f8) {
      let _this = this;
      (_this).f13 = f13;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      return;
    }
    fm0(p0, globalState) {
      let _this = this;
      return _dafny.Set.fromElements((_dafny.Set.fromElements(false)).Union(_dafny.Set.fromElements(false)), (_dafny.Set.fromElements(false)).Difference(_dafny.Set.fromElements(true, !(true))), _dafny.Set.fromElements(true, false, false));
    };
    fm1(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("nup"))).Union(function () {
        let _coll23 = new _dafny.Set();
        for (const _compr_23 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(620)), function (_282_i2) {
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(539)), function (_283_i3) {
            return new _dafny.CodePoint('h'.codePointAt(0));
          });
        })).Elements) {
          let _284_v0 = _compr_23;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(620)), function (_282_i2) {
            return _dafny.Seq.Create(_module.__default.abs(new BigNumber(539)), function (_283_i3) {
              return new _dafny.CodePoint('h'.codePointAt(0));
            });
          }), _284_v0)) {
            _coll23.add(_284_v0);
          }
        }
        return _coll23;
      }())).IsSubsetOf((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("jg"))).Union(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("q"), _dafny.Seq.UnicodeFromString("xbglft"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(739)), function (_285_i0) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-819)), function (_286_i1) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("gkesmfsv"))));
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f7 = _dafny.ZERO;
      this._f8 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    __ctor(f7, f8) {
      let _this = this;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      return;
    }
    fm18(globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jof"), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("hno"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(20)), function (_287_i0) {
        return new _dafny.CodePoint('k'.codePointAt(0));
      })));
    };
    fm19(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_module.D1.create_DC1(new _dafny.CodePoint('m'.codePointAt(0)))).dtor_cf1;
    };
    m6(globalState) {
      let _this = this;
      let r0 = _dafny.Set.Empty;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let _288_v0;
      let _init3 = function (_289_i0) {
        return ((_this).f7).isLessThan((_this).f7);
      };
      let _nw55 = Array((new BigNumber(6)).toNumber());
      for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw55.length); _i0_3++) {
        _nw55[_i0_3] = _init3(new BigNumber(_i0_3));
      }
      _288_v0 = _nw55;
      let _290_v1;
      _290_v1 = false;
      let _index48 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length));
      (_288_v0)[_index48] = _290_v1;
      let _291_v2;
      _291_v2 = _290_v1;
      let _292_v3;
      _292_v3 = _dafny.Seq.of(_290_v1, _290_v1, ((_290_v1) ? (_290_v1) : (_290_v1)));
      let _293_v4;
      _293_v4 = _dafny.Set.fromElements(new BigNumber(309));
      let _294_v5;
      _294_v5 = _module.D2.create_DC5(_290_v1, (_this).f7, new BigNumber((_293_v4).length), (_this).f7, _290_v1);
      let _pat_let_tv2 = _291_v2;
      let _pat_let_tv3 = _291_v2;
      let _index49 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length));
      let _rhs37 = _290_v1;
      let _rhs38 = (_292_v3)[_module.__default.safeIndex((_this).f8, new BigNumber((_292_v3).length))];
      let _rhs39 = function (_source7) {
        if (_source7.is_DC5) {
          let _295___mcc_h0 = (_source7).cf4;
          let _296___mcc_h1 = (_source7).cf5;
          let _297___mcc_h2 = (_source7).cf6;
          let _298___mcc_h3 = (_source7).cf7;
          let _299___mcc_h4 = (_source7).cf8;
          let _300_cf8 = _299___mcc_h4;
          let _301_cf7 = _298___mcc_h3;
          let _302_cf6 = _297___mcc_h2;
          let _303_cf5 = _296___mcc_h1;
          let _304_cf4 = _295___mcc_h0;
          return _pat_let_tv2;
        } else {
          let _305___mcc_h5 = (_source7).cf3;
          let _306_cf3 = _305___mcc_h5;
          return _pat_let_tv3;
        }
      }(_294_v5);
      let _rhs40 = (_291_v2);
      let _lhs38 = globalState;
      let _lhs39 = _288_v0;
      let _lhs40 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length));
      let _lhs41 = globalState;
      _lhs38.f4 = _rhs37;
      _lhs39[_lhs40] = _rhs38;
      _291_v2 = _rhs39;
      _lhs41.f4 = _rhs40;
      let _307_v6;
      _307_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,(_this).f7);
      let _308_v7;
      _308_v7 = _dafny.Seq.UnicodeFromString("tuyj");
      let _hi1 = new BigNumber(422);
      for (let _309_i1 = (((_307_v6).contains(_module.__default.fm20(_290_v1, (_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))], true, globalState))) ? ((_307_v6).get(_module.__default.fm20(_290_v1, (_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))], true, globalState))) : (new BigNumber((_308_v7).length))); _309_i1.isLessThan(_hi1); _309_i1 = _309_i1.plus(_dafny.ONE)) {
        let _310_v8;
        let _nw56 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
        _310_v8 = _nw56;
        let _index50 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_310_v8).length));
        (_310_v8)[_index50] = (_this).f7;
        let _311_v9;
        let _init4 = ((_312_v6) => function (_313_i2) {
          return _dafny.Seq.of((_dafny.ZERO).minus((_this).f7), (_this).f7, (((_312_v6).contains((_this).f8)) ? ((_312_v6).get((_this).f8)) : ((_this).f8)));
        })(_307_v6);
        let _nw57 = Array((new BigNumber(27)).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw57.length); _i0_4++) {
          _nw57[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _311_v9 = _nw57;
        let _index51 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_310_v8).length));
        let _rhs41 = (_this).f8;
        let _rhs42 = (_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))];
        let _rhs43 = _311_v9;
        let _rhs44 = (_294_v5).dtor_cf7;
        let _lhs42 = _310_v8;
        let _lhs43 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_310_v8).length));
        _lhs42[_lhs43] = _rhs41;
        _290_v1 = _rhs42;
        _311_v9 = _rhs43;
        r2 = _rhs44;
        let _314_v10;
        let _nw58 = new _module.C0();
        _nw58.__ctor(_310_v8, _309_i1, _309_i1);
        _314_v10 = _nw58;
        _314_v10 = _314_v10;
        let _315_v11;
        _315_v11 = new _dafny.CodePoint('y'.codePointAt(0));
        _315_v11 = _315_v11;
        let _316_v12;
        _316_v12 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('w'.codePointAt(0)),_290_v1);
        let _317_v13;
        _317_v13 = _dafny.Map.Empty.slice().updateUnsafe(_316_v12,(_this).f8);
        if ((_317_v13).contains(_316_v12)) {
          (globalState).f3 = (_292_v3)[_module.__default.safeIndex((_294_v5).dtor_cf5, new BigNumber((_292_v3).length))];
          let _index52 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_310_v8).length));
          let _rhs45 = _290_v1;
          let _rhs46 = (_310_v8)[_module.__default.safeIndex(new BigNumber(310), new BigNumber((_310_v8).length))];
          let _lhs44 = globalState;
          let _lhs45 = _310_v8;
          let _lhs46 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_310_v8).length));
          _lhs44.f4 = _rhs45;
          _lhs45[_lhs46] = _rhs46;
          let _318_v14;
          _318_v14 = _dafny.Map.Empty.slice().updateUnsafe(_291_v2,(_310_v8)[_module.__default.safeIndex(new BigNumber(310), new BigNumber((_310_v8).length))]);
          let _319_v15;
          _319_v15 = _dafny.Set.fromElements(_318_v14);
          let _320_v16;
          _320_v16 = _dafny.Seq.of(_293_v4);
          let _index53 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_310_v8).length));
          let _index54 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length));
          let _rhs47 = (_310_v8)[_module.__default.safeIndex(new BigNumber(310), new BigNumber((_310_v8).length))];
          let _rhs48 = _308_v7;
          let _rhs49 = _319_v15;
          let _rhs50 = _290_v1;
          let _rhs51 = (_320_v16)[_module.__default.safeIndex(((_this).f8).plus(new BigNumber(310)), new BigNumber((_320_v16).length))];
          let _lhs47 = _310_v8;
          let _lhs48 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_310_v8).length));
          let _lhs49 = _288_v0;
          let _lhs50 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length));
          _lhs47[_lhs48] = _rhs47;
          _308_v7 = _rhs48;
          _319_v15 = _rhs49;
          _lhs49[_lhs50] = _rhs50;
          _293_v4 = _rhs51;
          (globalState).f3 = (_309_i1).isLessThanOrEqualTo((_this).f8);
          let _321_v17;
          _321_v17 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f8).isLessThanOrEqualTo(_309_i1),(_310_v8)[_module.__default.safeIndex(new BigNumber(310), new BigNumber((_310_v8).length))]);
          _321_v17 = _321_v17;
        } else {
          _308_v7 = _dafny.Seq.Concat(_308_v7, _308_v7);
          let _322_v18;
          _322_v18 = _module.D1.create_DC1(_315_v11);
          let _323_v19;
          _323_v19 = _dafny.Seq.of(_322_v18, _322_v18, _module.D1.create_DC1(new _dafny.CodePoint('c'.codePointAt(0))));
          let _324_v20;
          let _nw59 = new _module.C0();
          _nw59.__ctor(_310_v8, new BigNumber((_module.__default.fm21(new BigNumber(715), new BigNumber((_323_v19).length), (_this).f7, _308_v7, globalState)).length), (_this).f8);
          _324_v20 = _nw59;
          let _325_v21;
          _325_v21 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_294_v5),_290_v1);
          let _326_v22;
          _326_v22 = _dafny.Seq.of(_294_v5);
          _325_v21 = (_325_v21).update(_326_v22, false);
          let _index55 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_310_v8).length));
          (_310_v8)[_index55] = _module.__default.safeDivisionInt((_310_v8)[_module.__default.safeIndex(new BigNumber(310), new BigNumber((_310_v8).length))], _module.__default.safeDivisionInt(_module.__default.fm20((_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))], (_290_v1), (_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))], globalState), new BigNumber(225)));
          let _327_v23;
          _327_v23 = _module.D1.create_DC1(_315_v11);
          let _328_v24;
          _328_v24 = _module.D1.create_DC3(_327_v23);
          let _329_v25;
          _329_v25 = _dafny.MultiSet.fromElements(_328_v24, _module.D1.create_DC3(_327_v23), _328_v24, _328_v24);
          _329_v25 = _module.__default.fm22(globalState);
        }
      }
      let _source8 = _module.D2.create_DC5((_292_v3)[_module.__default.safeIndex((_this).f7, new BigNumber((_292_v3).length))], (_this).f7, new BigNumber((_293_v4).length), (_this).f7, (_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))]);
      if (_source8.is_DC5) {
        let _330___mcc_h6 = (_source8).cf4;
        let _331___mcc_h7 = (_source8).cf5;
        let _332___mcc_h8 = (_source8).cf6;
        let _333___mcc_h9 = (_source8).cf7;
        let _334___mcc_h10 = (_source8).cf8;
        let _335_cf8 = _334___mcc_h10;
        let _336_cf7 = _333___mcc_h9;
        let _337_cf6 = _332___mcc_h8;
        let _338_cf5 = _331___mcc_h7;
        let _339_cf4 = _330___mcc_h6;
        let _340_v26;
        _340_v26 = new _dafny.CodePoint('q'.codePointAt(0));
        let _341_v27;
        _341_v27 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("xbuccmlho"), _module.__default.safeIndex((_this).f8, new BigNumber((_dafny.Seq.UnicodeFromString("xbuccmlho")).length)), _340_v26),(_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))]);
        let _342_v28;
        _342_v28 = _dafny.Map.Empty.slice().updateUnsafe((_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))],!(_341_v27).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(787)), ((_343_v26) => function (_344_i3) {
          return _343_v26;
        })(_340_v26)),_339_cf4)));
        _342_v28 = (_342_v28).update(_339_cf4, _335_cf8);
        let _345_v29;
        _345_v29 = _module.D1.create_DC3(_module.__default.fm23(new BigNumber((_dafny.Seq.UnicodeFromString("qf")).length), globalState));
        let _346_v30;
        _346_v30 = _module.D1.create_DC3(_345_v29);
        let _347_v31;
        _347_v31 = _dafny.Seq.of(_338_cf5);
        let _index56 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length));
        let _rhs52 = _346_v30;
        let _rhs53 = _290_v1;
        let _rhs54 = _339_cf4;
        let _rhs55 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-405)), ((_348_v26) => function (_349_i4) {
          return _348_v26;
        })(_340_v26)), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(602)), ((_350_v26) => function (_351_i5) {
          return _350_v26;
        })(_340_v26)), _308_v7)), _module.__default.safeIndex((_336_cf7).minus(new BigNumber((_347_v31).length)), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-405)), ((_352_v26) => function (_353_i4) {
          return _352_v26;
        })(_340_v26)), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(602)), ((_354_v26) => function (_355_i5) {
          return _354_v26;
        })(_340_v26)), _308_v7))).length)), _340_v26);
        let _lhs51 = _288_v0;
        let _lhs52 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length));
        _346_v30 = _rhs52;
        _339_cf4 = _rhs53;
        _lhs51[_lhs52] = _rhs54;
        _308_v7 = _rhs55;
        _339_cf4 = _module.__default.fm24((_336_cf7).isLessThan((_294_v5).dtor_cf5), _338_cf5, _dafny.Seq.of(new _dafny.CodePoint('c'.codePointAt(0))), globalState);
        let _index57 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length));
        let _rhs56 = (_this).f8;
        let _rhs57 = !(_339_cf4) || (_339_cf4);
        let _rhs58 = !(_290_v1);
        let _lhs53 = _288_v0;
        let _lhs54 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length));
        _338_cf5 = _rhs56;
        _lhs53[_lhs54] = _rhs57;
        _335_cf8 = _rhs58;
      } else {
        let _356___mcc_h11 = (_source8).cf3;
        let _357_cf3 = _356___mcc_h11;
        let _358_v32;
        let _init5 = function (_359_i6) {
          return (_359_i6).multipliedBy((_this).f7);
        };
        let _nw60 = Array((new BigNumber(6)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw60.length); _i0_5++) {
          _nw60[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _358_v32 = _nw60;
        let _index58 = _module.__default.safeIndex(new BigNumber(161), new BigNumber((_358_v32).length));
        (_358_v32)[_index58] = (_this).f8;
        let _360_v33;
        _360_v33 = _dafny.Set.fromElements((_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))]);
        let _361_v34;
        let _nw61 = new _module.C0();
        _nw61.__ctor(_358_v32, new BigNumber((_308_v7).length), (_this).f7);
        _361_v34 = _nw61;
        let _362_v35;
        _362_v35 = _dafny.Map.Empty.slice().updateUnsafe(_290_v1,((false) ? (_361_v34) : (_361_v34)));
        let _index59 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length));
        let _index60 = _module.__default.safeIndex(new BigNumber(161), new BigNumber((_358_v32).length));
        let _rhs59 = !(!((!(((_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))]) || (!((_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))])))) && (_290_v1)));
        let _rhs60 = ((_290_v1) ? ((_360_v33).Intersect(_360_v33)) : (_360_v33));
        let _rhs61 = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_362_v35).length)));
        let _rhs62 = ((_this).f7).plus((_this).f7);
        let _lhs55 = _288_v0;
        let _lhs56 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length));
        let _lhs57 = _358_v32;
        let _lhs58 = _module.__default.safeIndex(new BigNumber(161), new BigNumber((_358_v32).length));
        _lhs55[_lhs56] = _rhs59;
        r0 = _rhs60;
        _lhs57[_lhs58] = _rhs61;
        r2 = _rhs62;
        if (!((_dafny.ZERO).minus((_this).f7)).isEqualTo(((_this).f7).multipliedBy((_dafny.ZERO).minus((_this).f8)))) {
          let _index61 = _module.__default.safeIndex(new BigNumber(161), new BigNumber((_358_v32).length));
          (_358_v32)[_index61] = new BigNumber(117);
          let _363_v36;
          _363_v36 = _dafny.Seq.of(_357_cf3);
          let _364_v37;
          let _nw62 = Array((new BigNumber(11)).toNumber());
          _nw62[(_dafny.ZERO).toNumber()] = ((_357_cf3).update((_292_v3)[_module.__default.safeIndex((_this).f8, new BigNumber((_292_v3).length))], _module.__default.abs((_358_v32)[_module.__default.safeIndex(new BigNumber(161), new BigNumber((_358_v32).length))]))).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(!((_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))]), _290_v1)));
          _nw62[(_dafny.ONE).toNumber()] = _357_cf3;
          _nw62[(new BigNumber(2)).toNumber()] = _357_cf3;
          _nw62[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.FromArray(_292_v3);
          _nw62[(new BigNumber(4)).toNumber()] = _357_cf3;
          _nw62[(new BigNumber(5)).toNumber()] = _357_cf3;
          _nw62[(new BigNumber(6)).toNumber()] = _357_cf3;
          _nw62[(new BigNumber(7)).toNumber()] = (_357_cf3).Union(_357_cf3);
          _nw62[(new BigNumber(8)).toNumber()] = ((_363_v36)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f8), new BigNumber((_363_v36).length))]).Union(_357_cf3);
          _nw62[(new BigNumber(9)).toNumber()] = _357_cf3;
          _nw62[(new BigNumber(10)).toNumber()] = _357_cf3;
          _364_v37 = _nw62;
          let _index62 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_364_v37).length));
          (_364_v37)[_index62] = _357_cf3;
          let _index63 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_364_v37).length));
          (_364_v37)[_index63] = ((_357_cf3).Intersect(_dafny.MultiSet.fromElements(_290_v1))).Difference(_357_cf3);
          let _index64 = _module.__default.safeIndex(new BigNumber(161), new BigNumber((_358_v32).length));
          (_358_v32)[_index64] = ((_this).f8).multipliedBy(((_290_v1) ? ((_dafny.ZERO).minus((_358_v32)[_module.__default.safeIndex(new BigNumber(161), new BigNumber((_358_v32).length))])) : ((_358_v32)[_module.__default.safeIndex(new BigNumber(161), new BigNumber((_358_v32).length))])));
          let _365_v38;
          _365_v38 = new _dafny.CodePoint('u'.codePointAt(0));
          let _366_v39;
          _366_v39 = _dafny.Map.Empty.slice().updateUnsafe(_365_v38,(_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))]);
          _366_v39 = (_366_v39).update(_365_v38, (_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))]);
          let _index65 = _module.__default.safeIndex(new BigNumber(161), new BigNumber((_358_v32).length));
          (_358_v32)[_index65] = (_358_v32)[_module.__default.safeIndex(new BigNumber(161), new BigNumber((_358_v32).length))];
        } else {
          r1 = (((_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))]) ? (((_290_v1) ? ((_dafny.ZERO).minus((_this).f8)) : ((_this).f8))) : ((_358_v32)[_module.__default.safeIndex(new BigNumber(161), new BigNumber((_358_v32).length))]));
          let _367_v40;
          _367_v40 = _dafny.Seq.of(_module.D1.create_DC2());
          _367_v40 = _367_v40;
          let _368_v41;
          _368_v41 = _module.D2.create_DC4(_357_cf3);
          _368_v41 = _368_v41;
          let _369_v42;
          let _nw63 = new _module.C0();
          _nw63.__ctor(_358_v32, _module.__default.safeDivisionInt(new BigNumber(-810), (_358_v32)[_module.__default.safeIndex(new BigNumber(161), new BigNumber((_358_v32).length))]), ((_358_v32)[_module.__default.safeIndex(new BigNumber(161), new BigNumber((_358_v32).length))]).plus(new BigNumber((_308_v7).length)));
          _369_v42 = _nw63;
          let _370_v43;
          let _nw64 = Array((new BigNumber(14)).toNumber());
          _370_v43 = _nw64;
          let _index66 = _module.__default.safeIndex(new BigNumber(650), new BigNumber((_370_v43).length));
          (_370_v43)[_index66] = _361_v34;
          let _371_v44;
          _371_v44 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,_290_v1);
          let _372_v45;
          _372_v45 = _dafny.MultiSet.fromElements(_module.__default.fm20((_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))], _290_v1, (((_371_v44).contains((_this).f7)) ? ((_371_v44).get((_this).f7)) : ((_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))])), globalState), new BigNumber((_360_v33).length), (_this).f8);
          let _index67 = _module.__default.safeIndex(new BigNumber(650), new BigNumber((_370_v43).length));
          let _rhs63 = _369_v42;
          let _rhs64 = new BigNumber((_372_v45).cardinality());
          let _lhs59 = _370_v43;
          let _lhs60 = _module.__default.safeIndex(new BigNumber(650), new BigNumber((_370_v43).length));
          _lhs59[_lhs60] = _rhs63;
          r1 = _rhs64;
        }
        let _index68 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length));
        (_288_v0)[_index68] = (_292_v3)[_module.__default.safeIndex(((_290_v1) ? ((_358_v32)[_module.__default.safeIndex(new BigNumber(161), new BigNumber((_358_v32).length))]) : ((_358_v32)[_module.__default.safeIndex(new BigNumber(161), new BigNumber((_358_v32).length))])), new BigNumber((_292_v3).length))];
        let _373_v46;
        _373_v46 = _dafny.Set.fromElements(_291_v2);
        let _374_v47;
        _374_v47 = _dafny.Map.Empty.slice().updateUnsafe(_373_v46,_module.D2.create_DC5((_361_v34).fm1(_360_v33, (_358_v32)[_module.__default.safeIndex(new BigNumber(161), new BigNumber((_358_v32).length))], (_358_v32)[_module.__default.safeIndex(new BigNumber(161), new BigNumber((_358_v32).length))], (_this).f7, globalState), _module.__default.fm20(_290_v1, (_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))], (_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))], globalState), (_this).f8, new BigNumber(-418), _290_v1));
        let _index69 = _module.__default.safeIndex(new BigNumber(161), new BigNumber((_358_v32).length));
        (_358_v32)[_index69] = new BigNumber((_374_v47).length);
      }
      let _375_v48;
      _375_v48 = new _dafny.CodePoint('x'.codePointAt(0));
      let _376_v49;
      _376_v49 = _module.D1.create_DC1(_375_v48);
      let _377_v50;
      _377_v50 = _module.D1.create_DC3(_376_v49);
      let _source9 = _377_v50;
      if (_source9.is_DC2) {
        let _378_v51;
        _378_v51 = _dafny.Set.fromElements(_288_v0);
        if (((_378_v51).IsSubsetOf(_378_v51)) && (_290_v1)) {
          (globalState).f4 = _290_v1;
          let _379_v52;
          _379_v52 = _module.D1.create_DC1(((true) ? (_375_v48) : ((_this).fm19(_290_v1, (_294_v5).dtor_cf8, (_this).f8, (_this).f7, globalState))));
          _379_v52 = _379_v52;
          let _380_v53;
          let _nw65 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
          _380_v53 = _nw65;
          let _381_v54;
          _381_v54 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,_288_v0);
          let _382_v55;
          let _nw66 = new _module.C0();
          _nw66.__ctor(_380_v53, (_dafny.ZERO).minus((_this).f7), new BigNumber((_381_v54).length));
          _382_v55 = _nw66;
          let _383_v56;
          _383_v56 = _dafny.Map.Empty.slice().updateUnsafe((_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))],(new BigNumber(299)).isLessThanOrEqualTo(_module.__default.fm20(_290_v1, (_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))], (_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))], globalState)));
          _383_v56 = _module.__default.fm25(globalState);
          r2 = _module.__default.safeDivisionInt(((_290_v1) ? ((_this).f8) : ((_this).f7)), (_this).f8);
        } else {
          let _384_v57;
          _384_v57 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f8).multipliedBy(_module.__default.fm20(_290_v1, _290_v1, (_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))], globalState)),(_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))]);
          _384_v57 = (_384_v57).update((_this).f7, (_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))]);
          let _385_v58;
          let _nw67 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          _385_v58 = _nw67;
          let _index70 = _module.__default.safeIndex(new BigNumber(299), new BigNumber((_385_v58).length));
          (_385_v58)[_index70] = (_this).f8;
          let _index71 = _module.__default.safeIndex(new BigNumber(299), new BigNumber((_385_v58).length));
          (_385_v58)[_index71] = (_this).f7;
          let _386_v59;
          _386_v59 = _dafny.Seq.of((_385_v58)[_module.__default.safeIndex(new BigNumber(299), new BigNumber((_385_v58).length))]);
          let _387_v60;
          _387_v60 = _dafny.MultiSet.fromElements((_385_v58)[_module.__default.safeIndex(new BigNumber(299), new BigNumber((_385_v58).length))]);
          let _388_v61;
          _388_v61 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_386_v59, _module.__default.safeIndex(_module.__default.fm20(_module.__default.fm24(_290_v1, new BigNumber((_dafny.Seq.UnicodeFromString("trxrrblv")).length), _dafny.Seq.of(_375_v48, _375_v48), globalState), (_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))], _290_v1, globalState), new BigNumber((_386_v59).length)), (_this).f8),(_293_v4).IsProperSubsetOf((_module.__default.fm26(_387_v60, (_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))], globalState)).dtor_cf9));
          _388_v61 = _388_v61;
          r2 = (_this).f8;
          _290_v1 = (_292_v3)[_module.__default.safeIndex((_385_v58)[_module.__default.safeIndex(new BigNumber(299), new BigNumber((_385_v58).length))], new BigNumber((_292_v3).length))];
        }
        r1 = (_294_v5).dtor_cf5;
        (globalState).f4 = !(_290_v1);
        let _389_v62;
        _389_v62 = _dafny.Seq.of((_this).f7, (_this).f8);
        r2 = (new BigNumber(-476)).minus((_389_v62)[_module.__default.safeIndex(_module.__default.fm20((_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))], _290_v1, false, globalState), new BigNumber((_389_v62).length))]);
      } else if (_source9.is_DC1) {
        let _390___mcc_h12 = (_source9).cf1;
        let _391_cf1 = _390___mcc_h12;
        let _392_v63;
        let _init6 = function (_393_i7) {
          return (_393_i7).minus((_this).f7);
        };
        let _nw68 = Array((new BigNumber(27)).toNumber());
        for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw68.length); _i0_6++) {
          _nw68[_i0_6] = _init6(new BigNumber(_i0_6));
        }
        _392_v63 = _nw68;
        let _394_v64;
        _394_v64 = _dafny.Map.Empty.slice().updateUnsafe(_290_v1,(_dafny.ZERO).minus((_this).f7));
        let _index72 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_392_v63).length));
        (_392_v63)[_index72] = (new BigNumber((_394_v64).length)).multipliedBy((_294_v5).dtor_cf7);
        let _index73 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_392_v63).length));
        (_392_v63)[_index73] = new BigNumber(-885);
        let _395_v65;
        _395_v65 = _dafny.MultiSet.fromElements((_392_v63)[_module.__default.safeIndex(new BigNumber(300), new BigNumber((_392_v63).length))], (_this).f7);
        (globalState).f4 = (_395_v65).equals(_395_v65);
        let _396_v66;
        _396_v66 = _module.D3.create_DC7((_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))]);
        let _397_v67;
        _397_v67 = _dafny.Map.Empty.slice().updateUnsafe(_290_v1,_396_v66);
        _397_v67 = (_397_v67).update(((_290_v1) ? ((_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))]) : ((_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))])), _396_v66);
        if ((_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))]) {
          (globalState).f4 = (_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))];
          _391_cf1 = _391_cf1;
          (globalState).f4 = ((new BigNumber(957)).minus(new BigNumber(350))).isEqualTo((_this).f8);
          let _398_v68;
          let _nw69 = Array((new BigNumber(17)).toNumber());
          _nw69[(_dafny.ZERO).toNumber()] = _392_v63;
          _nw69[(_dafny.ONE).toNumber()] = _392_v63;
          _nw69[(new BigNumber(2)).toNumber()] = _392_v63;
          _nw69[(new BigNumber(3)).toNumber()] = _392_v63;
          _nw69[(new BigNumber(4)).toNumber()] = _392_v63;
          _nw69[(new BigNumber(5)).toNumber()] = _392_v63;
          _nw69[(new BigNumber(6)).toNumber()] = _392_v63;
          _nw69[(new BigNumber(7)).toNumber()] = _392_v63;
          _nw69[(new BigNumber(8)).toNumber()] = _392_v63;
          _nw69[(new BigNumber(9)).toNumber()] = _392_v63;
          _nw69[(new BigNumber(10)).toNumber()] = _392_v63;
          _nw69[(new BigNumber(11)).toNumber()] = _392_v63;
          _nw69[(new BigNumber(12)).toNumber()] = _392_v63;
          _nw69[(new BigNumber(13)).toNumber()] = _392_v63;
          _nw69[(new BigNumber(14)).toNumber()] = _392_v63;
          _nw69[(new BigNumber(15)).toNumber()] = _392_v63;
          _nw69[(new BigNumber(16)).toNumber()] = _392_v63;
          _398_v68 = _nw69;
          let _399_v69;
          _399_v69 = _dafny.Seq.of(_398_v68, _398_v68);
          let _400_v70;
          _400_v70 = _dafny.Map.Empty.slice().updateUnsafe((_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))],_290_v1);
          let _401_v71;
          let _nw70 = Array((new BigNumber(4)).toNumber());
          _nw70[(_dafny.ZERO).toNumber()] = _398_v68;
          _nw70[(_dafny.ONE).toNumber()] = (_399_v69)[_module.__default.safeIndex((_this).f8, new BigNumber((_399_v69).length))];
          _nw70[(new BigNumber(2)).toNumber()] = _398_v68;
          _nw70[(new BigNumber(3)).toNumber()] = (_399_v69)[_module.__default.safeIndex(new BigNumber((_400_v70).length), new BigNumber((_399_v69).length))];
          _401_v71 = _nw70;
          let _index74 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_401_v71).length));
          (_401_v71)[_index74] = _398_v68;
          let _402_v72;
          _402_v72 = _module.D4.create_DC8(_398_v68);
          let _index75 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_401_v71).length));
          (_401_v71)[_index75] = (_402_v72).dtor_cf11;
          let _403_v73;
          _403_v73 = _dafny.Map.Empty.slice().updateUnsafe((_392_v63)[_module.__default.safeIndex(new BigNumber(300), new BigNumber((_392_v63).length))],true);
          let _index76 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_392_v63).length));
          (_392_v63)[_index76] = _module.__default.safeModuloInt((new BigNumber((_dafny.Seq.of(_290_v1)).length)).minus(_module.__default.fm20(_290_v1, _290_v1, (_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))], globalState)), _module.__default.safeDivisionInt((_this).f8, new BigNumber((_403_v73).length)));
        } else {
          let _404_v74;
          let _nw71 = Array((new BigNumber(7)).toNumber());
          _nw71[(_dafny.ZERO).toNumber()] = _308_v7;
          _nw71[(_dafny.ONE).toNumber()] = _308_v7;
          _nw71[(new BigNumber(2)).toNumber()] = (_this).fm18(globalState);
          _nw71[(new BigNumber(3)).toNumber()] = _308_v7;
          _nw71[(new BigNumber(4)).toNumber()] = _308_v7;
          _nw71[(new BigNumber(5)).toNumber()] = (_this).fm18(globalState);
          _nw71[(new BigNumber(6)).toNumber()] = _308_v7;
          _404_v74 = _nw71;
          let _405_v75;
          _405_v75 = _dafny.Map.Empty.slice().updateUnsafe(_290_v1,_404_v74);
          _405_v75 = (_405_v75).update(_290_v1, _404_v74);
          let _406_v76;
          _406_v76 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_292_v3));
          let _407_v78;
          _407_v78 = _dafny.Map.Empty.slice().updateUnsafe(_290_v1,(_294_v5).dtor_cf8);
          let _408_v79;
          _408_v79 = _dafny.MultiSet.fromElements((((_407_v78).contains((_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))])) ? ((_407_v78).get((_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))])) : ((_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))])));
          let _409_v80;
          _409_v80 = _dafny.Set.fromElements(_408_v79);
          let _index77 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length));
          (_288_v0)[_index77] = !(((_module.__default.fm24(_290_v1, (_this).f8, _308_v7, globalState)) ? (function () {
            let _coll24 = new _dafny.Set();
            for (const _compr_24 of (_406_v76).Elements) {
              let _410_v77 = _compr_24;
              if (_dafny.Seq.contains(_406_v76, _410_v77)) {
                _coll24.add(_410_v77);
              }
            }
            return _coll24;
          }()) : (_module.__default.fm27(globalState)))).equals(_409_v80);
          (globalState).f3 = !_dafny.areEqual(_dafny.Seq.Concat(_292_v3, _292_v3), _292_v3);
          let _rhs65 = (_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))];
          let _rhs66 = (_this).f8;
          let _lhs61 = globalState;
          _lhs61.f4 = _rhs65;
          r1 = _rhs66;
          _307_v6 = ((_307_v6).update((_392_v63)[_module.__default.safeIndex(new BigNumber(300), new BigNumber((_392_v63).length))], (_dafny.ZERO).minus((_this).f7))).Merge(_307_v6);
        }
      } else {
        let _411___mcc_h13 = (_source9).cf2;
        let _412_cf2 = _411___mcc_h13;
        let _413_v81;
        _413_v81 = _dafny.MultiSet.fromElements((_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))]);
        let _414_v82;
        _414_v82 = _dafny.Seq.of(_413_v81, _413_v81, _dafny.MultiSet.fromElements(!(false), false));
        let _415_v83;
        _415_v83 = _dafny.MultiSet.fromElements((_this).f7);
        let _416_v84;
        let _nw72 = Array((new BigNumber(24)).toNumber());
        _nw72[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(970)), function (_417_i8) {
          return (_this).f7;
        })).length);
        _nw72[(_dafny.ONE).toNumber()] = (_this).f8;
        _nw72[(new BigNumber(2)).toNumber()] = (_this).f7;
        _nw72[(new BigNumber(3)).toNumber()] = (_this).f7;
        _nw72[(new BigNumber(4)).toNumber()] = (_this).f8;
        _nw72[(new BigNumber(5)).toNumber()] = new BigNumber(487);
        _nw72[(new BigNumber(6)).toNumber()] = (_this).f8;
        _nw72[(new BigNumber(7)).toNumber()] = (_this).f8;
        _nw72[(new BigNumber(8)).toNumber()] = new BigNumber((_308_v7).length);
        _nw72[(new BigNumber(9)).toNumber()] = (_this).f7;
        _nw72[(new BigNumber(10)).toNumber()] = (_this).f7;
        _nw72[(new BigNumber(11)).toNumber()] = new BigNumber(912);
        _nw72[(new BigNumber(12)).toNumber()] = new BigNumber((_dafny.Set.fromElements((_this).f8, (_294_v5).dtor_cf5, _module.__default.fm20(_290_v1, _290_v1, (_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))], globalState))).length);
        _nw72[(new BigNumber(13)).toNumber()] = (_this).f7;
        _nw72[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.Seq.update(_414_v82, _module.__default.safeIndex((_this).f8, new BigNumber((_414_v82).length)), _413_v81)).length);
        _nw72[(new BigNumber(15)).toNumber()] = (_this).f7;
        _nw72[(new BigNumber(16)).toNumber()] = new BigNumber((_307_v6).length);
        _nw72[(new BigNumber(17)).toNumber()] = (_this).f7;
        _nw72[(new BigNumber(18)).toNumber()] = new BigNumber((_dafny.Seq.update(_308_v7, _module.__default.safeIndex((_this).f8, new BigNumber((_308_v7).length)), _375_v48)).length);
        _nw72[(new BigNumber(19)).toNumber()] = new BigNumber((_415_v83).cardinality());
        _nw72[(new BigNumber(20)).toNumber()] = new BigNumber(609);
        _nw72[(new BigNumber(21)).toNumber()] = (_this).f8;
        _nw72[(new BigNumber(22)).toNumber()] = (_this).f8;
        _nw72[(new BigNumber(23)).toNumber()] = (_this).f8;
        _416_v84 = _nw72;
        let _418_v85;
        _418_v85 = _dafny.Map.Empty.slice().updateUnsafe(_290_v1,new _dafny.CodePoint('l'.codePointAt(0)));
        let _419_v86;
        let _nw73 = new _module.C0();
        _nw73.__ctor(_416_v84, (((_415_v83).contains((_this).f7)) ? ((_415_v83).get((_this).f7)) : (new BigNumber((_418_v85).length))), (_this).f8);
        _419_v86 = _nw73;
        let _index78 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length));
        let _index79 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length));
        let _rhs67 = _module.__default.fm24((_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))], (_this).f8, _308_v7, globalState);
        let _rhs68 = _290_v1;
        let _rhs69 = (_module.__default.fm20((_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))], (_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))], true, globalState)).isEqualTo((_this).f7);
        let _rhs70 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(90)), ((_420_v48) => function (_421_i9) {
          return _420_v48;
        })(_375_v48));
        let _lhs62 = _288_v0;
        let _lhs63 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length));
        let _lhs64 = _288_v0;
        let _lhs65 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length));
        let _lhs66 = globalState;
        _lhs62[_lhs63] = _rhs67;
        _lhs64[_lhs65] = _rhs68;
        _lhs66.f3 = _rhs69;
        _308_v7 = _rhs70;
        let _422_v87;
        let _init7 = ((_423_v1, _424_v0, _425_v83) => function (_426_i10) {
          return _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(_module.D1.create_DC2(), _module.D1.create_DC2(), _module.D1.create_DC2(), _module.D1.create_DC2(), _module.D1.create_DC2())).length), (_this).f8, new BigNumber((_dafny.Seq.of(_423_v1, (_424_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_424_v0).length))], _423_v1)).length), (_this).f8, new BigNumber((_425_v83).cardinality()));
        })(_290_v1, _288_v0, _415_v83);
        let _nw74 = Array((new BigNumber(25)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw74.length); _i0_7++) {
          _nw74[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _422_v87 = _nw74;
        let _427_v88;
        _427_v88 = _dafny.Seq.of(_288_v0);
        let _index80 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((_422_v87).length));
        (_422_v87)[_index80] = _module.__default.fm21(_module.__default.fm20((_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))], (_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))], true, globalState), (_this).f7, new BigNumber((_dafny.Seq.update(_427_v88, _module.__default.safeIndex((_this).f7, new BigNumber((_427_v88).length)), _288_v0)).length), _308_v7, globalState);
        let _428_v89;
        _428_v89 = _dafny.Map.Empty.slice().updateUnsafe((_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))],(_this).f8);
        let _429_v90;
        _429_v90 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_428_v89).length)));
        let _index81 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((_422_v87).length));
        (_422_v87)[_index81] = _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_292_v3).length), (_this).f7), _429_v90);
        if (_290_v1) {
          let _430_v91;
          _430_v91 = _module.D2.create_DC4((_module.__default.fm28((_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))], globalState)).Difference(_dafny.MultiSet.fromElements((_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))])));
          _430_v91 = _430_v91;
          let _arr7 = _419_v86.f13;
          let _index82 = _module.__default.safeIndex(new BigNumber(799), new BigNumber((_419_v86.f13).length));
          _arr7[_index82] = (_dafny.ZERO).minus((_this).f8);
          let _arr8 = _419_v86.f13;
          let _index83 = _module.__default.safeIndex(new BigNumber(799), new BigNumber((_419_v86.f13).length));
          _arr8[_index83] = new BigNumber((_dafny.Set.fromElements(_290_v1, (_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))])).length);
          let _431_v92;
          _431_v92 = _dafny.MultiSet.fromElements(_307_v6, _307_v6, _307_v6);
          let _index84 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length));
          (_288_v0)[_index84] = (_431_v92).IsProperSubsetOf(_431_v92);
          _413_v81 = _413_v81;
          let _432_v93;
          _432_v93 = _dafny.Map.Empty.slice().updateUnsafe((_419_v86.f13)[_module.__default.safeIndex(new BigNumber(799), new BigNumber((_419_v86.f13).length))],_290_v1);
          let _433_v94;
          let _nw75 = Array((new BigNumber(5)).toNumber()).fill([]);
          _433_v94 = _nw75;
          let _index85 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_433_v94).length));
          (_433_v94)[_index85] = _419_v86.f13;
          let _arr9 = _419_v86.f13;
          let _index86 = _module.__default.safeIndex(new BigNumber(799), new BigNumber((_419_v86.f13).length));
          let _index87 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_433_v94).length));
          let _rhs71 = (((_428_v89).contains(((_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))]) && (_290_v1))) ? ((_428_v89).get(((_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))]) && (_290_v1))) : (_module.__default.safeModuloInt((_this).f7, (_419_v86.f13)[_module.__default.safeIndex(new BigNumber(799), new BigNumber((_419_v86.f13).length))])));
          let _rhs72 = _432_v93;
          let _rhs73 = (_this).f7;
          let _rhs74 = _module.__default.safeModuloInt((_this).f8, (_this).f8);
          let _rhs75 = _419_v86.f13;
          let _lhs67 = _419_v86.f13;
          let _lhs68 = _module.__default.safeIndex(new BigNumber(799), new BigNumber((_419_v86.f13).length));
          let _lhs69 = _433_v94;
          let _lhs70 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_433_v94).length));
          _lhs67[_lhs68] = _rhs71;
          _432_v93 = _rhs72;
          r1 = _rhs73;
          r1 = _rhs74;
          _lhs69[_lhs70] = _rhs75;
        } else {
          r2 = (_module.__default.safeModuloInt((_this).f8, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_419_v86.f13,_290_v1)).length))).plus((_this).f8);
          let _434_v95;
          _434_v95 = _dafny.Map.Empty.slice().updateUnsafe((_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))],(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.update((_422_v87)[_module.__default.safeIndex(new BigNumber(487), new BigNumber((_422_v87).length))], _module.__default.safeIndex(new BigNumber(((_module.D5.create_DC11(_428_v89)).dtor_cf18).length), new BigNumber(((_422_v87)[_module.__default.safeIndex(new BigNumber(487), new BigNumber((_422_v87).length))]).length)), (_this).f7))).cardinality())).isLessThanOrEqualTo(new BigNumber((_dafny.Set.fromElements((_this).f7)).length)));
          let _435_v96;
          _435_v96 = _dafny.Seq.of(_434_v95, _434_v95, _434_v95, (_434_v95).Merge(_434_v95));
          _434_v95 = (_435_v96)[_module.__default.safeIndex((_this).f7, new BigNumber((_435_v96).length))];
          let _436_v97;
          let _nw76 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _436_v97 = _nw76;
          let _437_v98;
          _437_v98 = _dafny.MultiSet.fromElements(_375_v48);
          let _index88 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length));
          let _rhs76 = (_437_v98).IsProperSubsetOf(_437_v98);
          let _rhs77 = _436_v97;
          let _rhs78 = _290_v1;
          let _lhs71 = globalState;
          let _lhs72 = _288_v0;
          let _lhs73 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length));
          _lhs71.f4 = _rhs76;
          _436_v97 = _rhs77;
          _lhs72[_lhs73] = _rhs78;
          r2 = (_this).f8;
          r1 = (_this).f8;
        }
      }
      let _438_v99;
      _438_v99 = _module.D4.create_DC10((_this).f7, (_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))]);
      let _pat_let_tv4 = _308_v7;
      r2 = new BigNumber((function (_source10) {
        if (_source10.is_DC9) {
          let _439___mcc_h14 = (_source10).cf12;
          let _440___mcc_h15 = (_source10).cf13;
          let _441___mcc_h16 = (_source10).cf14;
          let _442___mcc_h17 = (_source10).cf15;
          let _443_cf15 = _442___mcc_h17;
          let _444_cf14 = _441___mcc_h16;
          let _445_cf13 = _440___mcc_h15;
          let _446_cf12 = _439___mcc_h14;
          return _dafny.Seq.UnicodeFromString("cradsw");
        } else if (_source10.is_DC10) {
          let _447___mcc_h18 = (_source10).cf16;
          let _448___mcc_h19 = (_source10).cf17;
          let _449_cf17 = _448___mcc_h19;
          let _450_cf16 = _447___mcc_h18;
          return _dafny.Seq.UnicodeFromString("bqqe");
        } else {
          let _451___mcc_h20 = (_source10).cf11;
          let _452_cf11 = _451___mcc_h20;
          return _pat_let_tv4;
        }
      }(_438_v99)).length);
      (globalState).f4 = ((((_288_v0)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_288_v0).length))]) ? (_290_v1) : ((_292_v3)[_module.__default.safeIndex((_this).f7, new BigNumber((_292_v3).length))]))) && (_290_v1);
      let _453_v100;
      _453_v100 = _dafny.Set.fromElements(_290_v1);
      r0 = (_453_v100).Intersect((_453_v100).Intersect(_453_v100));
      r1 = (_this).f8;
      let _454_v101;
      _454_v101 = _dafny.MultiSet.fromElements(_288_v0);
      r2 = ((_this).f8).plus(((false) ? (new BigNumber((_454_v101).cardinality())) : ((_this).f7)));
      return [r0, r1, r2];
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f7 = _dafny.ZERO;
      this._f8 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0, _module.T2, _module.T1];
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    __ctor(f7, f8) {
      let _this = this;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      return;
    }
    fm2(p0, globalState) {
      let _this = this;
      return _dafny.Seq.UnicodeFromString("ikasy");
    };
    fm0(p0, globalState) {
      let _this = this;
      return _dafny.Set.fromElements((_dafny.Set.fromElements(false, false, true, false)).Difference(_dafny.Set.fromElements(false)), (_dafny.Set.fromElements(false, true)).Difference(_dafny.Set.fromElements(true, false, true, true)), (_dafny.Set.fromElements(!(false), true)).Intersect(_dafny.Set.fromElements((false))));
    };
    fm1(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (((false) ? (false) : (true)));
    };
    fm14(globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qx"), _dafny.Seq.UnicodeFromString("cv"));
    };
    fm15(globalState) {
      let _this = this;
      return false;
    };
    m1(p0, globalState) {
      let _this = this;
      let _455_v0;
      let _init8 = function (_456_i0) {
        return ((_this).f8).isLessThan((_this).f8);
      };
      let _nw77 = Array((new BigNumber(19)).toNumber());
      for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw77.length); _i0_8++) {
        _nw77[_i0_8] = _init8(new BigNumber(_i0_8));
      }
      _455_v0 = _nw77;
      _455_v0 = _455_v0;
      let _457_v1;
      _457_v1 = false;
      let _458_v2;
      _458_v2 = _module.D2.create_DC5(_457_v1, (_this).f8, (_this).f7, (_this).f7, _457_v1);
      let _hi2 = (_458_v2).dtor_cf5;
      for (let _459_i1 = (_this).f7; _459_i1.isLessThan(_hi2); _459_i1 = _459_i1.plus(_dafny.ONE)) {
        let _460_v3;
        _460_v3 = _dafny.Seq.UnicodeFromString("rsl");
        _460_v3 = _dafny.Seq.Concat(p0, p0);
        let _461_v4;
        let _nw78 = Array((_dafny.ONE).toNumber()).fill([]);
        _461_v4 = _nw78;
        let _nw79 = Array((new BigNumber(10)).toNumber()).fill([]);
        _461_v4 = _nw79;
        let _462_v5;
        let _init9 = ((_463_i1) => function (_464_i2) {
          return (_464_i2).plus(_463_i1);
        })(_459_i1);
        let _nw80 = Array((new BigNumber(7)).toNumber());
        for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw80.length); _i0_9++) {
          _nw80[_i0_9] = _init9(new BigNumber(_i0_9));
        }
        _462_v5 = _nw80;
        let _465_v6;
        _465_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(375),(_this).f7);
        let _466_v7;
        let _nw81 = new _module.C0();
        _nw81.__ctor(_462_v5, new BigNumber((_dafny.Seq.Concat(_460_v3, p0)).length), new BigNumber((_465_v6).length));
        _466_v7 = _nw81;
        if ((new BigNumber(860)).isLessThan(new BigNumber(((_this).fm2(_457_v1, globalState)).length))) {
          let _467_v8;
          _467_v8 = new BigNumber(-417);
          _467_v8 = (new BigNumber(-602)).minus((_this).f8);
          _467_v8 = _467_v8;
          (globalState).f4 = _457_v1;
          (globalState).f3 = ((_457_v1) ? (_457_v1) : ((new BigNumber(-965)).isLessThanOrEqualTo(_459_i1)));
          let _index89 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_455_v0).length));
          (_455_v0)[_index89] = _457_v1;
          let _index90 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_455_v0).length));
          (_455_v0)[_index90] = !(_457_v1);
        } else {
          let _468_v9;
          _468_v9 = _dafny.MultiSet.fromElements((_this).f8);
          let _469_v10;
          _469_v10 = _dafny.Seq.of(_468_v9, _468_v9, _468_v9);
          let _470_v11;
          _470_v11 = _dafny.Map.Empty.slice().updateUnsafe(_455_v0,(_this).f7);
          let _471_v12;
          _471_v12 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_470_v11).length),_465_v6);
          let _472_v13;
          _472_v13 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_469_v10).length),(_471_v12).Merge(_471_v12));
          _472_v13 = (_472_v13).update(_module.__default.fm16(_457_v1, globalState), (_471_v12).update(_459_i1, _465_v6));
          let _473_v14;
          _473_v14 = _dafny.Seq.of((_this).f8);
          let _474_v15;
          _474_v15 = _dafny.Seq.of(_473_v14);
          let _475_v16;
          _475_v16 = _dafny.MultiSet.fromElements(_457_v1, _457_v1, !(_457_v1), _dafny.areEqual((_474_v15)[_module.__default.safeIndex((_458_v2).dtor_cf5, new BigNumber((_474_v15).length))], _473_v14));
          let _476_v17;
          _476_v17 = new _dafny.CodePoint('a'.codePointAt(0));
          let _477_v18;
          _477_v18 = _dafny.Map.Empty.slice().updateUnsafe(_457_v1,new BigNumber((_dafny.Seq.update(_460_v3, _module.__default.safeIndex((_dafny.ZERO).minus(_459_i1), new BigNumber((_460_v3).length)), _476_v17)).length));
          let _478_v19;
          _478_v19 = _dafny.Seq.of(true, !(!(_457_v1)), _457_v1);
          let _rhs79 = false;
          let _rhs80 = ((_475_v16).Union(_475_v16)).Union(_475_v16);
          let _rhs81 = (_477_v18).Merge(_477_v18);
          let _rhs82 = _478_v19;
          let _lhs74 = globalState;
          _lhs74.f3 = _rhs79;
          _475_v16 = _rhs80;
          _477_v18 = _rhs81;
          _478_v19 = _rhs82;
          let _479_v20;
          let _nw82 = new _module.C0();
          _nw82.__ctor(_466_v7.f13, (_this).f8, (_this).f7);
          _479_v20 = _nw82;
          let _480_v21;
          let _nw83 = Array((new BigNumber(2)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _480_v21 = _nw83;
          let _index91 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_480_v21).length));
          (_480_v21)[_index91] = (p0)[_module.__default.safeIndex(_module.__default.fm16(false, globalState), new BigNumber((p0).length))];
          let _index92 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_480_v21).length));
          (_480_v21)[_index92] = _476_v17;
          let _481_v22;
          let _nw84 = new _module.C0();
          _nw84.__ctor(((_457_v1) ? (_466_v7.f13) : (_462_v5)), _module.__default.safeModuloInt(_459_i1, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(286)), ((_482_v21) => function (_483_i3) {
            return (_482_v21)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_482_v21).length))];
          })(_480_v21))).length)), _module.__default.fm16(_457_v1, globalState));
          _481_v22 = _nw84;
        }
      }
      let _484_v24;
      _484_v24 = _dafny.MultiSet.fromElements(_457_v1, _457_v1);
      let _485_v25;
      _485_v25 = _dafny.Set.fromElements(new BigNumber((_484_v24).cardinality()), (_this).f8);
      let _486_v26;
      let _nw85 = Array((new BigNumber(2)).toNumber());
      _nw85[(_dafny.ZERO).toNumber()] = function () {
        let _coll25 = new _dafny.Set();
        for (const _compr_25 of _dafny.IntegerRange(new BigNumber(695), new BigNumber(-840))) {
          let _487_v23 = _compr_25;
          if (((new BigNumber(695)).isLessThanOrEqualTo(_487_v23)) && ((_487_v23).isLessThan(new BigNumber(-840)))) {
            _coll25.add((_487_v23).multipliedBy((_dafny.ZERO).minus((_this).f7)));
          }
        }
        return _coll25;
      }();
      _nw85[(_dafny.ONE).toNumber()] = _485_v25;
      _486_v26 = _nw85;
      let _488_v27;
      _488_v27 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f8))),_486_v26);
      let _489_v28;
      let _nw86 = Array((new BigNumber(11)).toNumber());
      _nw86[(_dafny.ZERO).toNumber()] = ((_457_v1) ? (_486_v26) : (_486_v26));
      _nw86[(_dafny.ONE).toNumber()] = (((_488_v27).contains((_this).f8)) ? ((_488_v27).get((_this).f8)) : (_486_v26));
      _nw86[(new BigNumber(2)).toNumber()] = _486_v26;
      _nw86[(new BigNumber(3)).toNumber()] = _486_v26;
      _nw86[(new BigNumber(4)).toNumber()] = _486_v26;
      _nw86[(new BigNumber(5)).toNumber()] = _486_v26;
      _nw86[(new BigNumber(6)).toNumber()] = _486_v26;
      _nw86[(new BigNumber(7)).toNumber()] = _486_v26;
      _nw86[(new BigNumber(8)).toNumber()] = _486_v26;
      _nw86[(new BigNumber(9)).toNumber()] = _486_v26;
      _nw86[(new BigNumber(10)).toNumber()] = _486_v26;
      _489_v28 = _nw86;
      let _index93 = _module.__default.safeIndex(new BigNumber(734), new BigNumber((_489_v28).length));
      (_489_v28)[_index93] = _486_v26;
      let _index94 = _module.__default.safeIndex(new BigNumber(734), new BigNumber((_489_v28).length));
      (_489_v28)[_index94] = _486_v26;
      let _490_v29;
      _490_v29 = _dafny.MultiSet.fromElements(new BigNumber(715), (_dafny.ZERO).minus((_this).f7));
      let _491_v30;
      _491_v30 = _dafny.Map.Empty.slice().updateUnsafe(_457_v1,_457_v1);
      let _492_v31;
      let _nw87 = Array((new BigNumber(25)).toNumber());
      _nw87[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_490_v29).cardinality()),_dafny.Seq.update(p0, _module.__default.safeIndex(new BigNumber((_484_v24).cardinality()), new BigNumber((p0).length)), new _dafny.CodePoint('y'.codePointAt(0))))).length);
      _nw87[(_dafny.ONE).toNumber()] = (_this).f8;
      _nw87[(new BigNumber(2)).toNumber()] = new BigNumber((_485_v25).length);
      _nw87[(new BigNumber(3)).toNumber()] = new BigNumber(285);
      _nw87[(new BigNumber(4)).toNumber()] = (_this).f7;
      _nw87[(new BigNumber(5)).toNumber()] = (_this).f8;
      _nw87[(new BigNumber(6)).toNumber()] = (_this).f8;
      _nw87[(new BigNumber(7)).toNumber()] = (_this).f8;
      _nw87[(new BigNumber(8)).toNumber()] = new BigNumber((p0).length);
      _nw87[(new BigNumber(9)).toNumber()] = (_this).f7;
      _nw87[(new BigNumber(10)).toNumber()] = (_this).f8;
      _nw87[(new BigNumber(11)).toNumber()] = (_this).f8;
      _nw87[(new BigNumber(12)).toNumber()] = new BigNumber(891);
      _nw87[(new BigNumber(13)).toNumber()] = _module.__default.fm16(_457_v1, globalState);
      _nw87[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus((_this).f7);
      _nw87[(new BigNumber(15)).toNumber()] = (_this).f8;
      _nw87[(new BigNumber(16)).toNumber()] = (_this).f7;
      _nw87[(new BigNumber(17)).toNumber()] = new BigNumber((p0).length);
      _nw87[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_491_v30).length));
      _nw87[(new BigNumber(19)).toNumber()] = (_this).f8;
      _nw87[(new BigNumber(20)).toNumber()] = (_this).f7;
      _nw87[(new BigNumber(21)).toNumber()] = (_this).f7;
      _nw87[(new BigNumber(22)).toNumber()] = (_this).f7;
      _nw87[(new BigNumber(23)).toNumber()] = (_this).f7;
      _nw87[(new BigNumber(24)).toNumber()] = new BigNumber((p0).length);
      _492_v31 = _nw87;
      let _493_v32;
      let _nw88 = new _module.C0();
      _nw88.__ctor(_492_v31, (_this).f7, (_this).f7);
      _493_v32 = _nw88;
      (globalState).f4 = _457_v1;
      let _494_v33;
      let _nw89 = new _module.C0();
      _nw89.__ctor(_493_v32.f13, (_dafny.ZERO).minus((_this).f8), new BigNumber(((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("trnq"), p0, p0, p0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(53)), function (_495_i4) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      }))).update(p0, _module.__default.abs((_this).f7))).cardinality()));
      _494_v33 = _nw89;
      return;
    }
    m2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = _dafny.ZERO;
      let r2 = false;
      let _496_v0;
      let _init10 = function (_497_i0) {
        return _module.__default.safeDivisionInt(_497_i0, (_this).f8);
      };
      let _nw90 = Array((new BigNumber(2)).toNumber());
      for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw90.length); _i0_10++) {
        _nw90[_i0_10] = _init10(new BigNumber(_i0_10));
      }
      _496_v0 = _nw90;
      let _498_v1;
      let _nw91 = new _module.C0();
      _nw91.__ctor(_496_v0, _module.__default.safeModuloInt(p0, (_this).f8), new BigNumber(765));
      _498_v1 = _nw91;
      (globalState).f4 = p1;
      let _499_v2;
      _499_v2 = _dafny.Seq.of(true, p3);
      let _500_v3;
      _500_v3 = _module.D2.create_DC4(_dafny.MultiSet.FromArray(_499_v2));
      let _source11 = _500_v3;
      if (_source11.is_DC5) {
        let _501___mcc_h0 = (_source11).cf4;
        let _502___mcc_h1 = (_source11).cf5;
        let _503___mcc_h2 = (_source11).cf6;
        let _504___mcc_h3 = (_source11).cf7;
        let _505___mcc_h4 = (_source11).cf8;
        let _506_cf8 = _505___mcc_h4;
        let _507_cf7 = _504___mcc_h3;
        let _508_cf6 = _503___mcc_h2;
        let _509_cf5 = _502___mcc_h1;
        let _510_cf4 = _501___mcc_h0;
        if (!(((_this).f8).isLessThan(_509_cf5))) {
          (globalState).f4 = _510_cf4;
          r2 = ((p3) ? (_506_cf8) : (((_506_cf8) ? (_506_cf8) : (false))));
          (globalState).f3 = p1;
          let _511_v4;
          _511_v4 = new _dafny.CodePoint('c'.codePointAt(0));
          let _512_v5;
          _512_v5 = _dafny.Map.Empty.slice().updateUnsafe(((true) ? (p3) : (false)),_511_v4);
          _512_v5 = (_512_v5).update(p1, new _dafny.CodePoint('e'.codePointAt(0)));
          let _513_v6;
          _513_v6 = _dafny.Seq.UnicodeFromString("ksccjtf");
          _513_v6 = _513_v6;
        } else {
          _508_cf6 = _module.__default.safeModuloInt(new BigNumber(((_this).fm14(globalState)).length), _507_cf7);
          let _514_v7;
          _514_v7 = new _dafny.CodePoint('g'.codePointAt(0));
          let _rhs83 = _514_v7;
          let _rhs84 = _499_v2;
          _514_v7 = _rhs83;
          r0 = _rhs84;
          let _515_v8;
          let _nw92 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _515_v8 = _nw92;
          let _516_v9;
          _516_v9 = _dafny.Seq.UnicodeFromString("ucud");
          let _index95 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_515_v8).length));
          (_515_v8)[_index95] = _dafny.Seq.update(_516_v9, _module.__default.safeIndex(new BigNumber(448), new BigNumber((_516_v9).length)), _514_v7);
          let _index96 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_515_v8).length));
          (_515_v8)[_index96] = _dafny.Seq.Concat(_516_v9, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-60)), ((_517_v7) => function (_518_i1) {
            return _517_v7;
          })(_514_v7)));
          r0 = _dafny.Seq.Concat(_dafny.Seq.of(p3), _dafny.Seq.Concat(_499_v2, _dafny.Seq.of(p3)));
          let _519_v10;
          let _nw93 = Array((new BigNumber(13)).toNumber()).fill(false);
          _519_v10 = _nw93;
          let _index97 = _module.__default.safeIndex(new BigNumber(272), new BigNumber((_519_v10).length));
          (_519_v10)[_index97] = _506_cf8;
          let _520_v11;
          _520_v11 = _dafny.Seq.of(_508_cf6);
          let _521_v12;
          _521_v12 = _dafny.Map.Empty.slice().updateUnsafe(_510_cf4,_520_v11);
          let _index98 = _module.__default.safeIndex(new BigNumber(272), new BigNumber((_519_v10).length));
          let _rhs85 = p3;
          let _rhs86 = p1;
          let _rhs87 = ((p1) ? (_508_cf6) : (_508_cf6));
          let _rhs88 = _520_v11;
          let _rhs89 = _module.__default.fm17(_514_v7, _510_cf4, globalState);
          let _lhs75 = _519_v10;
          let _lhs76 = _module.__default.safeIndex(new BigNumber(272), new BigNumber((_519_v10).length));
          _lhs75[_lhs76] = _rhs85;
          r2 = _rhs86;
          _508_cf6 = _rhs87;
          _520_v11 = _rhs88;
          _521_v12 = _rhs89;
        }
        if (p1) {
          let _522_v13;
          _522_v13 = _dafny.Seq.UnicodeFromString("v");
          _522_v13 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-528)), function (_523_i2) {
            return new _dafny.CodePoint('d'.codePointAt(0));
          }), _522_v13);
          let _524_v14;
          _524_v14 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,_dafny.Seq.UnicodeFromString("lb"));
          _524_v14 = (_524_v14).update(_508_cf6, _522_v13);
          let _525_v15;
          let _nw94 = new _module.C1();
          _nw94.__ctor((_this).f7, (_this).f8);
          _525_v15 = _nw94;
          _525_v15 = ((p1) ? (_525_v15) : (_525_v15));
          let _526_v16;
          _526_v16 = _dafny.Set.fromElements(_506_cf8);
          let _527_v17;
          _527_v17 = _dafny.Map.Empty.slice().updateUnsafe((_525_v15).f7,new BigNumber(833));
          (globalState).f3 = (_498_v1).fm1(_526_v16, (_this).f8, (_dafny.ZERO).minus(new BigNumber((_527_v17).length)), (_525_v15).f8, globalState);
          let _528_v18;
          _528_v18 = _dafny.Map.Empty.slice().updateUnsafe(p3,(_525_v15).f8);
          _528_v18 = (_528_v18).update(p1, (new BigNumber((function () {
            let _coll26 = new _dafny.Map();
            for (const _compr_26 of _dafny.IntegerRange(new BigNumber(437), new BigNumber(-935))) {
              let _529_v19 = _compr_26;
              if (((new BigNumber(437)).isLessThanOrEqualTo(_529_v19)) && ((_529_v19).isLessThan(new BigNumber(-935)))) {
                _coll26.push([_module.__default.safeModuloInt(_529_v19, _509_cf5),p3]);
              }
            }
            return _coll26;
          }()).length)).minus(p0));
        } else {
          let _530_v20;
          _530_v20 = _dafny.Map.Empty.slice().updateUnsafe(!(p3),(_509_cf5).isLessThan(p2));
          _530_v20 = (_530_v20).update(_506_cf8, _506_cf8);
          (globalState).f4 = (_module.__default.safeModuloInt(new BigNumber(915), _507_cf7)).isLessThan(_508_cf6);
          let _531_v21;
          _531_v21 = new _dafny.CodePoint('v'.codePointAt(0));
          let _532_v22;
          _532_v22 = _dafny.Seq.of(_531_v21);
          let _arr10 = _498_v1.f13;
          let _index99 = _module.__default.safeIndex(new BigNumber(904), new BigNumber((_498_v1.f13).length));
          _arr10[_index99] = new BigNumber((_dafny.Seq.Concat(_532_v22, _dafny.Seq.of(new _dafny.CodePoint('j'.codePointAt(0))))).length);
          let _533_v23;
          _533_v23 = _dafny.Map.Empty.slice().updateUnsafe(!(_506_cf8),p0);
          let _534_v24;
          _534_v24 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(p3,new BigNumber((_532_v22).length)));
          let _arr11 = _498_v1.f13;
          let _index100 = _module.__default.safeIndex(new BigNumber(904), new BigNumber((_498_v1.f13).length));
          _arr11[_index100] = ((_506_cf8) ? ((((_533_v23).contains(p3)) ? ((_533_v23).get(p3)) : (new BigNumber((_534_v24).cardinality())))) : (((_this).f7).minus(p2)));
          _530_v20 = (_530_v20).update(false, (new BigNumber((_532_v22).length)).isEqualTo(_507_cf7));
          let _535_v25;
          _535_v25 = _dafny.Seq.of(_509_cf5, _507_cf7);
          let _536_v26;
          _536_v26 = _dafny.MultiSet.fromElements(_535_v25);
          _536_v26 = _dafny.MultiSet.fromElements(_535_v25, _535_v25);
        }
        _509_cf5 = _507_cf7;
        (globalState).f4 = _510_cf4;
      } else {
        let _537___mcc_h5 = (_source11).cf3;
        let _538_cf3 = _537___mcc_h5;
        r1 = (_this).f8;
        let _539_v27;
        _539_v27 = new _dafny.CodePoint('y'.codePointAt(0));
        let _540_v28;
        _540_v28 = _dafny.Set.fromElements(_539_v27, _539_v27, new _dafny.CodePoint('d'.codePointAt(0)));
        let _541_v29;
        _541_v29 = _dafny.Seq.of(_540_v28);
        (globalState).f4 = (new BigNumber((_541_v29).length)).isLessThanOrEqualTo(p2);
        let _542_v30;
        let _nw95 = Array((new BigNumber(18)).toNumber()).fill(false);
        _542_v30 = _nw95;
        let _543_v31;
        _543_v31 = _dafny.Seq.of(_542_v30);
        let _544_v32;
        _544_v32 = _module.D4.create_DC9(p0, _543_v31, (_this).f7, p0);
        let _545_v33;
        let _nw96 = Array((new BigNumber(3)).toNumber()).fill([]);
        _545_v33 = _nw96;
        let _546_v34;
        _546_v34 = _module.D4.create_DC8(_545_v33);
        let _547_v35;
        _547_v35 = _dafny.Map.Empty.slice().updateUnsafe(p1,_546_v34);
        let _pat_let_tv5 = p0;
        let _pat_let_tv6 = _547_v35;
        let _source12 = function (_pat_let0_0) {
          return function (_551_dt__update__tmp_h1) {
            return function (_pat_let4_0) {
              return function (_552_dt__update_hcf6_h0) {
                return _module.D2.create_DC5((_551_dt__update__tmp_h1).dtor_cf4, (_551_dt__update__tmp_h1).dtor_cf5, _552_dt__update_hcf6_h0, (_551_dt__update__tmp_h1).dtor_cf7, (_551_dt__update__tmp_h1).dtor_cf8);
              }(_pat_let4_0);
            }((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_pat_let_tv6).length))));
          }(_pat_let0_0);
        }(function (_pat_let1_0) {
          return function (_548_dt__update__tmp_h0) {
            return function (_pat_let2_0) {
              return function (_549_dt__update_hcf5_h0) {
                return function (_pat_let3_0) {
                  return function (_550_dt__update_hcf8_h0) {
                    return _module.D2.create_DC5((_548_dt__update__tmp_h0).dtor_cf4, _549_dt__update_hcf5_h0, (_548_dt__update__tmp_h0).dtor_cf6, (_548_dt__update__tmp_h0).dtor_cf7, _550_dt__update_hcf8_h0);
                  }(_pat_let3_0);
                }(false);
              }(_pat_let2_0);
            }(_pat_let_tv5);
          }(_pat_let1_0);
        }(_module.D2.create_DC5(true, (_this).f8, new BigNumber(475), (_544_v32).dtor_cf14, p3)));
        if (_source12.is_DC5) {
          let _553___mcc_h6 = (_source12).cf4;
          let _554___mcc_h7 = (_source12).cf5;
          let _555___mcc_h8 = (_source12).cf6;
          let _556___mcc_h9 = (_source12).cf7;
          let _557___mcc_h10 = (_source12).cf8;
          let _558_cf8 = _557___mcc_h10;
          let _559_cf7 = _556___mcc_h9;
          let _560_cf6 = _555___mcc_h8;
          let _561_cf5 = _554___mcc_h7;
          let _562_cf4 = _553___mcc_h6;
          let _563_v36;
          _563_v36 = _dafny.Map.Empty.slice().updateUnsafe(!(_558_cf8),_498_v1.f13);
          let _564_v37;
          _564_v37 = _dafny.Set.fromElements(new BigNumber((_563_v36).length));
          let _565_v38;
          _565_v38 = _module.D3.create_DC6((_564_v37).Union(_564_v37));
          _565_v38 = _565_v38;
          (globalState).f4 = p3;
          let _566_v39;
          _566_v39 = _dafny.Set.fromElements(p1);
          let _index101 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_542_v30).length));
          (_542_v30)[_index101] = true;
          let _567_v40;
          _567_v40 = _dafny.Seq.of(_566_v39, _566_v39, _dafny.Set.fromElements((_499_v2)[_module.__default.safeIndex((_this).f7, new BigNumber((_499_v2).length))]));
          let _index102 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_542_v30).length));
          let _rhs90 = (_567_v40)[_module.__default.safeIndex((new BigNumber((_499_v2).length)).plus(_559_cf7), new BigNumber((_567_v40).length))];
          let _rhs91 = ((p3) ? (((_this).f8).isEqualTo(p0)) : (p1));
          let _lhs77 = _542_v30;
          let _lhs78 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_542_v30).length));
          _566_v39 = _rhs90;
          _lhs77[_lhs78] = _rhs91;
          _560_cf6 = (_this).f8;
        } else {
          let _568___mcc_h11 = (_source12).cf3;
          let _569_cf3 = _568___mcc_h11;
          (globalState).f4 = (p2).isLessThan((_this).f7);
          r1 = (_dafny.ZERO).minus(new BigNumber(-380));
          let _570_v41;
          let _nw97 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Set.Empty);
          _570_v41 = _nw97;
          let _571_v44;
          _571_v44 = _dafny.Set.fromElements(p0);
          let _index103 = _module.__default.safeIndex(new BigNumber(757), new BigNumber((_570_v41).length));
          (_570_v41)[_index103] = (function () {
            let _coll27 = new _dafny.Set();
            for (const _compr_27 of (function () {
              let _coll28 = new _dafny.Set();
              for (const _compr_28 of _dafny.IntegerRange(new BigNumber(-205), new BigNumber(46))) {
                let _572_v42 = _compr_28;
                if (((new BigNumber(-205)).isLessThanOrEqualTo(_572_v42)) && ((_572_v42).isLessThan(new BigNumber(46)))) {
                  _coll28.add(_module.__default.safeDivisionInt(_572_v42, new BigNumber((_dafny.Seq.UnicodeFromString("vudgnl")).length)));
                }
              }
              return _coll28;
            }()).Elements) {
              let _573_v43 = _compr_27;
              if ((function () {
                let _coll29 = new _dafny.Set();
                for (const _compr_29 of _dafny.IntegerRange(new BigNumber(-205), new BigNumber(46))) {
                  let _574_v42 = _compr_29;
                  if (((new BigNumber(-205)).isLessThanOrEqualTo(_574_v42)) && ((_574_v42).isLessThan(new BigNumber(46)))) {
                    _coll29.add(_module.__default.safeDivisionInt(_574_v42, new BigNumber((_dafny.Seq.UnicodeFromString("vudgnl")).length)));
                  }
                }
                return _coll29;
              }()).contains(_573_v43)) {
                _coll27.add(_module.__default.safeDivisionInt(_573_v43, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(432)), function (_575_i3) {
                  return new _dafny.CodePoint('o'.codePointAt(0));
                })).length)));
              }
            }
            return _coll27;
          }()).Difference(_571_v44);
          let _576_v46;
          _576_v46 = _module.D3.create_DC6(_dafny.Set.fromElements(p2));
          let _pat_let_tv7 = p2;
          let _pat_let_tv8 = _543_v31;
          let _index104 = _module.__default.safeIndex(new BigNumber(757), new BigNumber((_570_v41).length));
          let _rhs92 = (_540_v28).IsSubsetOf((_540_v28).Intersect(function () {
            let _coll30 = new _dafny.Set();
            for (const _compr_30 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(986)), ((_577_v27) => function (_578_i4) {
              return _577_v27;
            })(_539_v27))).Elements) {
              let _579_v45 = _compr_30;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(986)), ((_580_v27) => function (_578_i4) {
                return _580_v27;
              })(_539_v27)), _579_v45)) {
                _coll30.add(_579_v45);
              }
            }
            return _coll30;
          }()));
          let _rhs93 = _499_v2;
          let _rhs94 = (_this).f7;
          let _rhs95 = function (_pat_let5_0) {
            return function (_581_dt__update__tmp_h2) {
              return function (_pat_let6_0) {
                return function (_582_dt__update_hcf14_h0) {
                  return _module.D4.create_DC9((_581_dt__update__tmp_h2).dtor_cf12, (_581_dt__update__tmp_h2).dtor_cf13, _582_dt__update_hcf14_h0, (_581_dt__update__tmp_h2).dtor_cf15);
                }(_pat_let6_0);
              }((_pat_let_tv7).minus(new BigNumber((_pat_let_tv8).length)));
            }(_pat_let5_0);
          }(_544_v32);
          let _rhs96 = (_576_v46).dtor_cf9;
          let _lhs79 = _570_v41;
          let _lhs80 = _module.__default.safeIndex(new BigNumber(757), new BigNumber((_570_v41).length));
          r2 = _rhs92;
          r0 = _rhs93;
          r1 = _rhs94;
          _544_v32 = _rhs95;
          _lhs79[_lhs80] = _rhs96;
          _542_v30 = _542_v30;
        }
        let _583_v47;
        _583_v47 = _dafny.MultiSet.fromElements(p0, p2);
        let _584_v48;
        let _nw98 = new _module.C1();
        _nw98.__ctor((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.fm16(true, globalState))), (_dafny.ZERO).minus(new BigNumber((_module.__default.fm29(p1, _583_v47, globalState)).length)));
        _584_v48 = _nw98;
      }
      let _585_v49;
      _585_v49 = _dafny.MultiSet.fromElements(p1);
      let _586_v50;
      _586_v50 = _dafny.Seq.of(_module.__default.fm20(p3, !(false), p3, globalState), (((_585_v49).contains(p3)) ? ((_585_v49).get(p3)) : ((_dafny.ZERO).minus((_this).f7))));
      let _hi3 = p0;
      for (let _587_i5 = new BigNumber((((p3) ? (_586_v50) : (_586_v50))).length); _587_i5.isLessThan(_hi3); _587_i5 = _587_i5.plus(_dafny.ONE)) {
        (globalState).f3 = p1;
        let _588_v51;
        _588_v51 = _dafny.Seq.UnicodeFromString("nvldbus");
        let _589_v52;
        _589_v52 = _dafny.Map.Empty.slice().updateUnsafe(p3,new BigNumber((_588_v51).length));
        let _590_v53;
        _590_v53 = _dafny.MultiSet.fromElements(p0);
        _589_v52 = (_589_v52).update(p3, (((_590_v53).contains(p2)) ? ((_590_v53).get(p2)) : (p0)));
        let _591_v54;
        let _init11 = ((_592_p3) => function (_593_i6) {
          return _592_p3;
        })(p3);
        let _nw99 = Array((new BigNumber(13)).toNumber());
        for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw99.length); _i0_11++) {
          _nw99[_i0_11] = _init11(new BigNumber(_i0_11));
        }
        _591_v54 = _nw99;
        let _594_v55;
        _594_v55 = _module.D4.create_DC9(p2, _dafny.Seq.of(_591_v54, _591_v54, _591_v54, _591_v54, _591_v54), new BigNumber(688), (_this).f8);
        let _arr12 = _498_v1.f13;
        let _index105 = _module.__default.safeIndex(new BigNumber(503), new BigNumber((_498_v1.f13).length));
        _arr12[_index105] = _module.__default.safeDivisionInt((((_589_v52).contains(p1)) ? ((_589_v52).get(p1)) : (p2)), (_594_v55).dtor_cf14);
        let _arr13 = _498_v1.f13;
        let _index106 = _module.__default.safeIndex(new BigNumber(503), new BigNumber((_498_v1.f13).length));
        _arr13[_index106] = (p0).multipliedBy((_this).f7);
        let _595_v56;
        _595_v56 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_588_v51).length),_496_v0);
        let _596_v57;
        _596_v57 = _dafny.Map.Empty.slice().updateUnsafe(_595_v56,_590_v53);
        _596_v57 = (_596_v57).update(_595_v56, _dafny.MultiSet.FromArray(((p1) ? (_586_v50) : (_dafny.Seq.of(new BigNumber((_590_v53).cardinality()), (_this).f7, new BigNumber(-717), new BigNumber(620), new BigNumber(742))))));
      }
      r1 = _module.__default.fm16((_this).fm15(globalState), globalState);
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_496_v0).length))) {
        let _597_i7 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_597_i7)) && ((_597_i7).isLessThan(new BigNumber((_496_v0).length))))) {
          (_496_v0)[(_597_i7)] = _module.__default.safeModuloInt(_597_i7, (_this).f8);
        }
      }
      let _598_v58;
      _598_v58 = _dafny.Set.fromElements(p3);
      r0 = _dafny.Seq.update(_499_v2, _module.__default.safeIndex((_dafny.ZERO).minus((_this).f7), new BigNumber((_499_v2).length)), !((_this).f7).isEqualTo(new BigNumber((_598_v58).length)));
      let _599_v59;
      _599_v59 = _dafny.Set.fromElements(p0);
      let _600_v60;
      _600_v60 = _module.D3.create_DC6(_599_v59);
      let _pat_let_tv9 = p0;
      let _pat_let_tv10 = p2;
      r1 = function (_source13) {
        if (_source13.is_DC7) {
          let _601___mcc_h12 = (_source13).cf10;
          let _602_cf10 = _601___mcc_h12;
          if (_602_cf10) {
            return _pat_let_tv9;
          } else {
            return new BigNumber((_dafny.MultiSet.fromElements(_pat_let_tv10, new BigNumber(-808), (_this).f8, new BigNumber((_dafny.Seq.UnicodeFromString("fjemwxaw")).length), (_this).f8)).cardinality());
          }
        } else {
          let _603___mcc_h13 = (_source13).cf9;
          let _604_cf9 = _603___mcc_h13;
          return new BigNumber(332);
        }
      }(_600_v60);
      let _605_v61;
      _605_v61 = _dafny.Seq.UnicodeFromString("rtyih");
      r2 = !(_module.__default.fm20(false, p3, p1, globalState)).isEqualTo(new BigNumber((_605_v61).length));
      return [r0, r1, r2];
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f7 = _dafny.ZERO;
      this._f8 = _dafny.ZERO;
      this._f12 = false;
    }
    _parentTraits() {
      return [_module.T2, _module.T0, _module.T1];
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    __ctor(f12, f7, f8) {
      let _this = this;
      (_this)._f12 = f12;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      return;
    }
    fm2(p0, globalState) {
      let _this = this;
      return _dafny.Seq.UnicodeFromString("skp");
    };
    fm0(p0, globalState) {
      let _this = this;
      let _source14 = (_this).f12;
      let _606___mcc_h0 = _source14;
      let _607_cf0 = _606___mcc_h0;
      if ((_this).f12) {
        return _dafny.Set.fromElements(_dafny.Set.fromElements(false, _607_cf0));
      } else {
        return _dafny.Set.fromElements(_dafny.Set.fromElements(_607_cf0), _dafny.Set.fromElements((_this).f12));
      }
    };
    fm1(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f8, new BigNumber((_dafny.Seq.UnicodeFromString("xutwsx")).length), (_this).f8))).cardinality())).minus((_this).f8)).isLessThanOrEqualTo(_module.__default.safeModuloInt(new BigNumber((_dafny.Set.fromElements((_this).f12, (_this).f12, (_this).f12, (_this).f12, !((_this).f12))).length), new BigNumber((function () {
        let _coll31 = new _dafny.Map();
        for (const _compr_31 of _dafny.IntegerRange(new BigNumber(234), new BigNumber(342))) {
          let _608_v0 = _compr_31;
          if (((new BigNumber(234)).isLessThanOrEqualTo(_608_v0)) && ((_608_v0).isLessThan(new BigNumber(342)))) {
            _coll31.push([(_608_v0).minus(new BigNumber(906)),(_this).f12]);
          }
        }
        return _coll31;
      }()).length)));
    };
    fm12(globalState) {
      let _this = this;
      return (_this).f12;
    };
    fm13(p0, globalState) {
      let _this = this;
      return (_this).f8;
    };
    m1(p0, globalState) {
      let _this = this;
      let _609_v0;
      _609_v0 = _dafny.Seq.of((_this).f12, !((_this).f12), (_this).f12);
      let _hi4 = (_dafny.ZERO).minus((_this).f7);
      for (let _610_i0 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_609_v0, _609_v0)).length)); _610_i0.isLessThan(_hi4); _610_i0 = _610_i0.plus(_dafny.ONE)) {
        let _611_v1;
        _611_v1 = (_this).f12;
        if (!((new BigNumber((_dafny.MultiSet.fromElements(_611_v1)).cardinality())).isLessThan(_610_i0))) {
          let _612_v2;
          let _init12 = ((_613_i0) => function (_614_i1) {
            return _module.__default.safeModuloInt(_614_i1, _613_i0);
          })(_610_i0);
          let _nw100 = Array((new BigNumber(22)).toNumber());
          for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw100.length); _i0_12++) {
            _nw100[_i0_12] = _init12(new BigNumber(_i0_12));
          }
          _612_v2 = _nw100;
          let _615_v3;
          _615_v3 = _dafny.MultiSet.fromElements((_this).f7);
          let _616_v4;
          _616_v4 = _dafny.Map.Empty.slice().updateUnsafe(_612_v2,_615_v3);
          let _617_v5;
          let _init13 = function (_618_i2) {
            return !(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(735)), function (_619_i3) {
              return new _dafny.CodePoint('b'.codePointAt(0));
            }))).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('e'.codePointAt(0))));
          };
          let _nw101 = Array((new BigNumber(29)).toNumber());
          for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw101.length); _i0_13++) {
            _nw101[_i0_13] = _init13(new BigNumber(_i0_13));
          }
          _617_v5 = _nw101;
          let _index107 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_617_v5).length));
          (_617_v5)[_index107] = (_this).f12;
          let _620_v6;
          _620_v6 = _dafny.Set.fromElements((_this).f7, (_this).f8);
          let _621_v7;
          _621_v7 = _dafny.Seq.of(new BigNumber((_620_v6).length));
          let _index108 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_617_v5).length));
          let _rhs97 = (_616_v4).Merge(_616_v4);
          let _rhs98 = !_dafny.areEqual(_dafny.Seq.Concat(_621_v7, _dafny.Seq.of((_this).f7)), _dafny.Seq.Concat(_621_v7, _621_v7));
          let _rhs99 = (_this).f12;
          let _rhs100 = (_this).f12;
          let _lhs81 = globalState;
          let _lhs82 = globalState;
          let _lhs83 = _617_v5;
          let _lhs84 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_617_v5).length));
          _616_v4 = _rhs97;
          _lhs81.f3 = _rhs98;
          _lhs82.f4 = _rhs99;
          _lhs83[_lhs84] = _rhs100;
          let _622_v8;
          _622_v8 = new _dafny.CodePoint('r'.codePointAt(0));
          let _623_v9;
          _623_v9 = _module.D1.create_DC1((_module.D1.create_DC1(new _dafny.CodePoint('j'.codePointAt(0)))).dtor_cf1);
          _622_v8 = (_623_v9).dtor_cf1;
          (globalState).f3 = (_617_v5)[_module.__default.safeIndex(new BigNumber(259), new BigNumber((_617_v5).length))];
          (globalState).f3 = (_609_v0)[_module.__default.safeIndex(_610_i0, new BigNumber((_609_v0).length))];
          let _index109 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_617_v5).length));
          (_617_v5)[_index109] = (_this).f12;
        } else {
          (globalState).f4 = ((_this).f12) && ((((_this).f12) ? ((_this).f12) : ((_this).f12)));
          let _624_v10;
          _624_v10 = _dafny.MultiSet.fromElements(_609_v0);
          (globalState).f3 = (_624_v10).IsProperSubsetOf((_624_v10).Difference(_624_v10));
          let _625_v11;
          let _init14 = function (_626_i4) {
            return (_this).f12;
          };
          let _nw102 = Array((new BigNumber(28)).toNumber());
          for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw102.length); _i0_14++) {
            _nw102[_i0_14] = _init14(new BigNumber(_i0_14));
          }
          _625_v11 = _nw102;
          _625_v11 = _625_v11;
          let _627_v12;
          _627_v12 = _dafny.Seq.of(_609_v0);
          _627_v12 = _627_v12;
          let _628_v13;
          _628_v13 = new BigNumber(658);
          _628_v13 = ((_this).f7).plus(_module.__default.safeModuloInt((_this).f8, (_this).f7));
        }
        let _629_v14;
        _629_v14 = new _dafny.CodePoint('g'.codePointAt(0));
        _629_v14 = _629_v14;
        (globalState).f4 = (_this).f12;
        let _630_v15;
        _630_v15 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,_module.__default.safeModuloInt(_610_i0, new BigNumber((_dafny.MultiSet.fromElements((_this).f12)).cardinality())));
        _630_v15 = _630_v15;
      }
      let _631_v16;
      _631_v16 = _dafny.MultiSet.fromElements((_this).f8);
      let _632_v17;
      let _633_v18;
      let _out2;
      let _out3;
      let _outcollector1 = _module.__default.m0((_this).f7, (_this).fm1(_dafny.Set.fromElements((_this).f12, (_this).f12), new BigNumber(-562), (_this).f7, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber(893))).length),(_this).f8)).length), globalState), (_631_v16).IsSubsetOf(_631_v16), globalState);
      _out2 = _outcollector1[0];
      _out3 = _outcollector1[1];
      _632_v17 = _out2;
      _633_v18 = _out3;
      let _634_v19;
      _634_v19 = new _dafny.CodePoint('b'.codePointAt(0));
      let _635_v20;
      _635_v20 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm12(globalState),_634_v19);
      _635_v20 = _635_v20;
      let _636_i5;
      _636_i5 = _dafny.ZERO;
      L3: {
        while ((_this).f12) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_636_i5)) {
              break L3;
            }
            _636_i5 = (_636_i5).plus(_dafny.ONE);
            let _637_v21;
            _637_v21 = _dafny.Seq.UnicodeFromString("xronwep");
            let _rhs101 = p0;
            let _rhs102 = (_this).f12;
            _637_v21 = _rhs101;
            _633_v18 = _rhs102;
            _633_v18 = (_this).f12;
            let _638_v22;
            _638_v22 = new BigNumber(771);
            let _639_v23;
            let _nw103 = Array((_dafny.ONE).toNumber()).fill(false);
            _639_v23 = _nw103;
            let _640_v24;
            _640_v24 = _module.D1.create_DC1(_634_v19);
            let _641_v25;
            _641_v25 = _dafny.MultiSet.fromElements(_640_v24);
            let _642_v26;
            _642_v26 = _dafny.Seq.of(_641_v25);
            let _rhs103 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(643), new BigNumber(663)));
            let _rhs104 = _639_v23;
            let _rhs105 = _634_v19;
            let _rhs106 = _dafny.Seq.Concat(p0, _637_v21);
            let _rhs107 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(763)), ((_643_v24) => function (_644_i6) {
              return _dafny.MultiSet.FromArray(_dafny.Seq.of(_643_v24, _643_v24));
            })(_640_v24)), _642_v26);
            _638_v22 = _rhs103;
            _639_v23 = _rhs104;
            _634_v19 = _rhs105;
            _637_v21 = _rhs106;
            _642_v26 = _rhs107;
            let _645_v27;
            _645_v27 = _dafny.Map.Empty.slice().updateUnsafe(_633_v18,_633_v18);
            let _646_v28;
            _646_v28 = _dafny.Map.Empty.slice().updateUnsafe((_645_v27).Merge(_645_v27),_639_v23);
            _646_v28 = (_646_v28).update((_645_v27).Merge(_645_v27), _639_v23);
          }
        }
      }
      let _647_v29;
      _647_v29 = _dafny.MultiSet.fromElements(false);
      let _648_v30;
      _648_v30 = _dafny.Seq.of(new BigNumber((_647_v29).cardinality()));
      let _hi5 = ((_this).f7).multipliedBy((_this).f7);
      for (let _649_i7 = (_648_v30)[_module.__default.safeIndex((_this).f8, new BigNumber((_648_v30).length))]; _649_i7.isLessThan(_hi5); _649_i7 = _649_i7.plus(_dafny.ONE)) {
        let _650_v31;
        let _651_v32;
        let _out4;
        let _out5;
        let _outcollector2 = _module.__default.m0(((_this).f7).plus((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("vvjdbcbj"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-435)), ((_652_v19) => function (_653_i8) {
          return _652_v19;
        })(_634_v19)))).cardinality()))), _633_v18, (_633_v18) === (true), globalState);
        _out4 = _outcollector2[0];
        _out5 = _outcollector2[1];
        _650_v31 = _out4;
        _651_v32 = _out5;
        (globalState).f3 = (_651_v32);
        _609_v0 = _dafny.Seq.Concat(_609_v0, _dafny.Seq.Concat(_609_v0, _609_v0));
        let _rhs108 = _648_v30;
        let _rhs109 = (_this).f12;
        let _lhs85 = globalState;
        _648_v30 = _rhs108;
        _lhs85.f3 = _rhs109;
      }
      let _654_v33;
      _654_v33 = _dafny.Seq.of(((_633_v18) ? (_609_v0) : (_609_v0)));
      _609_v0 = (_654_v33)[_module.__default.safeIndex(_module.__default.safeDivisionInt((_this).f8, (_this).f7), new BigNumber((_654_v33).length))];
      return;
    }
    m2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = _dafny.ZERO;
      let r2 = false;
      let _655_v0;
      _655_v0 = _dafny.Seq.UnicodeFromString("dwwkgew");
      r1 = _module.__default.safeDivisionInt(new BigNumber(267), new BigNumber((_655_v0).length));
      let _656_v1;
      _656_v1 = new _dafny.CodePoint('m'.codePointAt(0));
      let _657_v2;
      let _init15 = ((_658_p2) => function (_659_i0) {
        return ((_this).f8).isLessThan((_dafny.ZERO).minus(_658_p2));
      })(p2);
      let _nw104 = Array((new BigNumber(20)).toNumber());
      for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw104.length); _i0_15++) {
        _nw104[_i0_15] = _init15(new BigNumber(_i0_15));
      }
      _657_v2 = _nw104;
      let _rhs110 = new _dafny.CodePoint('y'.codePointAt(0));
      let _rhs111 = _657_v2;
      let _rhs112 = (_dafny.ZERO).minus(new BigNumber((_655_v0).length));
      let _rhs113 = new BigNumber(255);
      _656_v1 = _rhs110;
      _657_v2 = _rhs111;
      r1 = _rhs112;
      r1 = _rhs113;
      r2 = p3;
      let _660_v3;
      _660_v3 = _dafny.Seq.of((_this).f12, true, p1);
      let _661_v4;
      _661_v4 = _dafny.MultiSet.fromElements(!(p1), (_660_v3)[_module.__default.safeIndex(p2, new BigNumber((_660_v3).length))]);
      let _662_v5;
      _662_v5 = _module.D2.create_DC4(_661_v4);
      let _pat_let_tv11 = _661_v4;
      _661_v4 = (function (_pat_let7_0) {
        return function (_663_dt__update__tmp_h0) {
          return function (_pat_let8_0) {
            return function (_664_dt__update_hcf3_h0) {
              return _module.D2.create_DC4(_664_dt__update_hcf3_h0);
            }(_pat_let8_0);
          }(_pat_let_tv11);
        }(_pat_let7_0);
      }(_662_v5)).dtor_cf3;
      let _hi6 = (_dafny.ZERO).minus((_this).f7);
      for (let _665_i1 = p2; _665_i1.isLessThan(_hi6); _665_i1 = _665_i1.plus(_dafny.ONE)) {
        r1 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_module.__default.safeModuloInt(p0, new BigNumber(-574)), new BigNumber(444)));
        let _666_v6;
        _666_v6 = (_this).f12;
        let _source15 = _666_v6;
        let _667___mcc_h0 = _source15;
        let _668_cf0 = _667___mcc_h0;
        r1 = p0;
        r1 = (_this).f8;
        let _669_v7;
        _669_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,(_this).fm13(_655_v0, globalState));
        r1 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_669_v7).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(p1), _660_v3)).length))));
        let _670_v8;
        let _nw105 = new _module.C2();
        _nw105.__ctor(p2, p2);
        _670_v8 = _nw105;
        _670_v8 = _670_v8;
        let _671_v9;
        _671_v9 = _module.D5.create_DC12(!((_this).fm12(globalState)));
        let _672_v10;
        _672_v10 = _dafny.MultiSet.fromElements(new BigNumber(-703), p0, new BigNumber((_655_v0).length));
        let _673_v11;
        _673_v11 = _dafny.Seq.of(_665_i1);
        let _rhs114 = ((_this).f7).isLessThan(new BigNumber((_module.__default.fm30(p3, globalState)).length));
        let _rhs115 = !((p3) || ((_this).f12));
        let _rhs116 = _671_v9;
        let _rhs117 = (_672_v10).Difference((_672_v10).Union(_dafny.MultiSet.FromArray(_673_v11)));
        let _lhs86 = globalState;
        r2 = _rhs114;
        _lhs86.f3 = _rhs115;
        _671_v9 = _rhs116;
        _672_v10 = _rhs117;
        (globalState).f4 = !(((p1) ? (false) : (!(p1))));
      }
      let _674_v12;
      let _init16 = ((_675_v3, _676_p1) => function (_677_i3) {
        return _module.__default.safeModuloInt(_677_i3, new BigNumber((_dafny.Seq.update(_675_v3, _module.__default.safeIndex((_this).f7, new BigNumber((_675_v3).length)), _676_p1)).length));
      })(_660_v3, p1);
      let _nw106 = Array((new BigNumber(24)).toNumber());
      for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw106.length); _i0_16++) {
        _nw106[_i0_16] = _init16(new BigNumber(_i0_16));
      }
      _674_v12 = _nw106;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_674_v12).length))) {
        let _678_i2 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_678_i2)) && ((_678_i2).isLessThan(new BigNumber((_674_v12).length))))) {
          (_674_v12)[(_678_i2)] = _module.__default.safeModuloInt(_678_i2, (((_this).f12) ? (p2) : ((_dafny.ZERO).minus((_this).f8))));
        }
      }
      r0 = _660_v3;
      r1 = (p2).multipliedBy((_this).f7);
      r2 = p3;
      return [r0, r1, r2];
    }
    get f12() {
      let _this = this;
      return _this._f12;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f7 = _dafny.ZERO;
      this._f8 = _dafny.ZERO;
      this._f11 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    __ctor(f11, f7, f8) {
      let _this = this;
      (_this)._f11 = f11;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      return;
    }
    fm0(p0, globalState) {
      let _this = this;
      return ((_dafny.Set.fromElements(_dafny.Set.fromElements((_this).f11))).Union(_dafny.Set.fromElements(_dafny.Set.fromElements((_this).f11, (_this).f11), _dafny.Set.fromElements((_this).f11, false, false, (_this).f11)))).Union((_dafny.Set.fromElements(_dafny.Set.fromElements((_this).f11))).Union(function () {
        let _coll32 = new _dafny.Set();
        for (const _compr_32 of (function () {
          let _coll33 = new _dafny.Set();
          for (const _compr_33 of (_dafny.MultiSet.fromElements(_dafny.Set.fromElements((_this).f11), _dafny.Set.fromElements((_this).f11, (_this).f11, (_this).f11))).Elements) {
            let _679_v0 = _compr_33;
            if ((_dafny.MultiSet.fromElements(_dafny.Set.fromElements((_this).f11), _dafny.Set.fromElements((_this).f11, (_this).f11, (_this).f11))).contains(_679_v0)) {
              _coll33.add(_679_v0);
            }
          }
          return _coll33;
        }()).Elements) {
          let _680_v1 = _compr_32;
          if ((function () {
            let _coll34 = new _dafny.Set();
            for (const _compr_34 of (_dafny.MultiSet.fromElements(_dafny.Set.fromElements((_this).f11), _dafny.Set.fromElements((_this).f11, (_this).f11, (_this).f11))).Elements) {
              let _681_v0 = _compr_34;
              if ((_dafny.MultiSet.fromElements(_dafny.Set.fromElements((_this).f11), _dafny.Set.fromElements((_this).f11, (_this).f11, (_this).f11))).contains(_681_v0)) {
                _coll34.add(_681_v0);
              }
            }
            return _coll34;
          }()).contains(_680_v1)) {
            _coll32.add(_680_v1);
          }
        }
        return _coll32;
      }()));
    };
    fm1(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(480)), function (_682_i0) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      })).length)).isLessThan((_this).f7);
    };
    fm11(globalState) {
      let _this = this;
      if ((_this).f11) {
        return (_this).f8;
      } else {
        return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of((_this).f7, (_this).f7), _dafny.Seq.Create(_module.__default.abs(new BigNumber(702)), function (_683_i0) {
          return new BigNumber(453);
        }))).length);
      }
    };
    m5(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = false;
      let r3 = false;
      r1 = p0;
      let _684_v0;
      let _nw107 = new _module.C2();
      _nw107.__ctor(new BigNumber(537), new BigNumber(-183));
      _684_v0 = _nw107;
      let _685_v1;
      let _nw108 = new _module.C2();
      _nw108.__ctor(p2, _module.__default.safeDivisionInt((_this).f7, (_this).f7));
      _685_v1 = _nw108;
      let _686_i0;
      _686_i0 = _dafny.ZERO;
      L4: {
        while ((((_this).f11) ? (p0) : (p0))) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_686_i0)) {
              break L4;
            }
            _686_i0 = (_686_i0).plus(_dafny.ONE);
            let _687_v2;
            _687_v2 = new BigNumber(241);
            _687_v2 = p2;
            let _688_v3;
            let _nw109 = new _module.C1();
            _nw109.__ctor(_module.__default.safeModuloInt((_this).f7, new BigNumber(83)), _module.__default.safeModuloInt((_this).f7, _module.__default.fm16(true, globalState)));
            _688_v3 = _nw109;
            let _689_v4;
            let _nw110 = Array((new BigNumber(9)).toNumber()).fill([]);
            _689_v4 = _nw110;
            _689_v4 = _689_v4;
            let _690_v5;
            _690_v5 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
            let _691_v6;
            _691_v6 = _dafny.Map.Empty.slice().updateUnsafe(_690_v5,_dafny.Seq.of(p0, (_this).f11));
            _687_v2 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(p2), _module.__default.safeDivisionInt(_687_v2, new BigNumber((_691_v6).length)));
          }
        }
      }
      let _source16 = _module.__default.fm26(_dafny.MultiSet.fromElements(p2), (_this).f11, globalState);
      if (_source16.is_DC7) {
        let _692___mcc_h0 = (_source16).cf10;
        let _693_cf10 = _692___mcc_h0;
        let _694_v7;
        _694_v7 = new BigNumber(184);
        let _695_v8;
        _695_v8 = _dafny.Seq.of((_this).f11, false);
        let _696_v9;
        _696_v9 = _dafny.Seq.of(_695_v8, _695_v8);
        _694_v7 = (_module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f8), _694_v7)).multipliedBy(new BigNumber(((_696_v9)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("lhbhuxmc")).length), new BigNumber((_696_v9).length))]).length));
        let _697_v10;
        _697_v10 = _dafny.MultiSet.fromElements((_694_v7).multipliedBy(p3));
        let _698_v11;
        _698_v11 = _dafny.Seq.of(p1);
        _694_v7 = (_dafny.ZERO).minus((((_697_v10).contains(p2)) ? ((_697_v10).get(p2)) : (new BigNumber((_dafny.Seq.Concat(_698_v11, _698_v11)).length))));
        _698_v11 = _698_v11;
        let _699_v12;
        _699_v12 = _dafny.Seq.UnicodeFromString("md");
        let _700_v13;
        let _nw111 = Array((new BigNumber(10)).toNumber());
        _nw111[(_dafny.ZERO).toNumber()] = (_this).f8;
        _nw111[(_dafny.ONE).toNumber()] = p1;
        _nw111[(new BigNumber(2)).toNumber()] = (_this).f8;
        _nw111[(new BigNumber(3)).toNumber()] = new BigNumber(-318);
        _nw111[(new BigNumber(4)).toNumber()] = p2;
        _nw111[(new BigNumber(5)).toNumber()] = _694_v7;
        _nw111[(new BigNumber(6)).toNumber()] = new BigNumber((_699_v12).length);
        _nw111[(new BigNumber(7)).toNumber()] = (_698_v11)[_module.__default.safeIndex(_694_v7, new BigNumber((_698_v11).length))];
        _nw111[(new BigNumber(8)).toNumber()] = (_this).f8;
        _nw111[(new BigNumber(9)).toNumber()] = p3;
        _700_v13 = _nw111;
        let _701_v14;
        let _nw112 = Array((_dafny.ONE).toNumber());
        _nw112[(_dafny.ZERO).toNumber()] = _700_v13;
        _701_v14 = _nw112;
        let _index110 = _module.__default.safeIndex(new BigNumber(304), new BigNumber((_701_v14).length));
        (_701_v14)[_index110] = _700_v13;
        let _index111 = _module.__default.safeIndex(new BigNumber(304), new BigNumber((_701_v14).length));
        (_701_v14)[_index111] = _700_v13;
      } else {
        let _702___mcc_h1 = (_source16).cf9;
        let _703_cf9 = _702___mcc_h1;
        let _704_v15;
        _704_v15 = _dafny.Seq.UnicodeFromString("u");
        let _705_v16;
        _705_v16 = _dafny.Set.fromElements(_704_v15);
        let _706_v17;
        let _init17 = ((_707_p0) => function (_708_i1) {
          return _707_p0;
        })(p0);
        let _nw113 = Array((new BigNumber(25)).toNumber());
        for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw113.length); _i0_17++) {
          _nw113[_i0_17] = _init17(new BigNumber(_i0_17));
        }
        _706_v17 = _nw113;
        let _709_v18;
        _709_v18 = _dafny.Map.Empty.slice().updateUnsafe(_704_v15,_706_v17);
        let _710_v19;
        _710_v19 = new BigNumber(-109);
        let _711_v20;
        _711_v20 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("ekpn"),_706_v17), _709_v18);
        let _712_v21;
        _712_v21 = _dafny.Set.fromElements(p0);
        let _713_v22;
        _713_v22 = _dafny.MultiSet.fromElements((_this).f11);
        let _rhs118 = (_705_v16).Difference(_705_v16);
        let _rhs119 = (_711_v20)[_module.__default.safeIndex(new BigNumber((_712_v21).length), new BigNumber((_711_v20).length))];
        let _rhs120 = new BigNumber(((_713_v22).Intersect(_713_v22)).cardinality());
        _705_v16 = _rhs118;
        _709_v18 = _rhs119;
        _710_v19 = _rhs120;
        r3 = !((_this).f11);
        let _714_v23;
        _714_v23 = _dafny.Seq.of(p0, p0);
        let _715_v24;
        _715_v24 = _dafny.Seq.of(new BigNumber(80), (_this).f7, p1);
        (globalState).f4 = _dafny.areEqual(_dafny.Seq.update(_dafny.Seq.Concat(_714_v23, _714_v23), _module.__default.safeIndex(new BigNumber((_715_v24).length), new BigNumber((_dafny.Seq.Concat(_714_v23, _714_v23)).length)), (_this).f11), _714_v23);
        _714_v23 = _dafny.Seq.update(_714_v23, _module.__default.safeIndex(new BigNumber(607), new BigNumber((_714_v23).length)), p0);
      }
      let _716_v25;
      _716_v25 = new BigNumber(321);
      _716_v25 = p2;
      r0 = p0;
      r1 = (_this).f11;
      r2 = !(p0) || ((_this).f11);
      r3 = !((_this).f11);
      return [r0, r1, r2, r3];
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f7 = _dafny.ZERO;
      this._f8 = _dafny.ZERO;
      this._f9 = _dafny.Seq.UnicodeFromString("");
      this._f10 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    __ctor(f9, f10, f7, f8) {
      let _this = this;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      return;
    }
    fm0(p0, globalState) {
      let _this = this;
      return ((_dafny.Set.fromElements(_dafny.Set.fromElements(false, false))).Difference(_dafny.Set.fromElements(_dafny.Set.fromElements(true, true), _dafny.Set.fromElements(false, true)))).Union(_dafny.Set.fromElements(_dafny.Set.fromElements(false), _dafny.Set.fromElements(true, false, !(true), false, true)));
    };
    fm1(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (false) && (!_dafny.areEqual(_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of((_this).f7)).length), (_this).f10, (_this).f8, new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f10, (_this).f10))).cardinality()), new BigNumber(((_this).f9).length)), _dafny.Set.fromElements((_this).f10)), _dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(64)))));
    };
    fm6(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f8;
    };
    fm7(p0, globalState) {
      let _this = this;
      return ((_this).f10).multipliedBy((_this).f7);
    };
    m4(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _717_v0;
      let _init18 = function (_718_i0) {
        return (_718_i0).multipliedBy((_this).f7);
      };
      let _nw114 = Array((new BigNumber(11)).toNumber());
      for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw114.length); _i0_18++) {
        _nw114[_i0_18] = _init18(new BigNumber(_i0_18));
      }
      _717_v0 = _nw114;
      let _index112 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_717_v0).length));
      (_717_v0)[_index112] = (_this).f7;
      let _719_v1;
      _719_v1 = false;
      let _index113 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_717_v0).length));
      (_717_v0)[_index113] = (_this).fm6(new _dafny.CodePoint('a'.codePointAt(0)), _module.__default.fm8(_719_v1, _dafny.Seq.of((_this).f8, p0, p0), _719_v1, globalState), _719_v1, globalState);
      let _720_v2;
      _720_v2 = _dafny.Seq.of(_719_v1, _719_v1, _719_v1);
      let _721_v3;
      _721_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_719_v1);
      let _722_v4;
      _722_v4 = new _dafny.CodePoint('g'.codePointAt(0));
      let _723_v5;
      _723_v5 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(((_this).f9).length));
      let _724_v6;
      _724_v6 = _719_v1;
      let _725_v7;
      _725_v7 = _dafny.Set.fromElements((_this).f8, new BigNumber(((_this).f9).length));
      let _726_v8;
      _726_v8 = _dafny.Map.Empty.slice().updateUnsafe(_719_v1,false);
      let _727_v9;
      _727_v9 = _dafny.MultiSet.fromElements(_module.__default.fm9(_719_v1, globalState));
      let _728_v10;
      _728_v10 = _dafny.Set.fromElements(_719_v1);
      let _729_v11;
      _729_v11 = _dafny.MultiSet.fromElements(_719_v1, _719_v1);
      let _730_v12;
      _730_v12 = _dafny.Seq.of((_this).f7);
      let _731_v13;
      _731_v13 = _dafny.Map.Empty.slice().updateUnsafe(((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_730_v12).length)))).update(p0, _module.__default.abs((_this).f7))).update(p0, _module.__default.abs((_this).f8)),_719_v1);
      let _732_v14;
      let _nw115 = Array((new BigNumber(27)).toNumber());
      _nw115[(_dafny.ZERO).toNumber()] = _719_v1;
      _nw115[(_dafny.ONE).toNumber()] = _719_v1;
      _nw115[(new BigNumber(2)).toNumber()] = (_720_v2)[_module.__default.safeIndex((_this).f8, new BigNumber((_720_v2).length))];
      _nw115[(new BigNumber(3)).toNumber()] = (((_721_v3).contains(_dafny.Seq.UnicodeFromString("twaaoc"))) ? ((_721_v3).get(_dafny.Seq.UnicodeFromString("twaaoc"))) : (_719_v1));
      _nw115[(new BigNumber(4)).toNumber()] = _719_v1;
      _nw115[(new BigNumber(5)).toNumber()] = _719_v1;
      _nw115[(new BigNumber(6)).toNumber()] = _719_v1;
      _nw115[(new BigNumber(7)).toNumber()] = _719_v1;
      _nw115[(new BigNumber(8)).toNumber()] = _719_v1;
      _nw115[(new BigNumber(9)).toNumber()] = ((_dafny.ZERO).minus((_this).fm6(_722_v4, (_this).f9, _719_v1, globalState))).isLessThan((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(_719_v1)).cardinality())));
      _nw115[(new BigNumber(10)).toNumber()] = !(_719_v1);
      _nw115[(new BigNumber(11)).toNumber()] = _719_v1;
      _nw115[(new BigNumber(12)).toNumber()] = !(_723_v5).contains(new BigNumber((_720_v2).length));
      _nw115[(new BigNumber(13)).toNumber()] = (_724_v6);
      _nw115[(new BigNumber(14)).toNumber()] = _719_v1;
      _nw115[(new BigNumber(15)).toNumber()] = _719_v1;
      _nw115[(new BigNumber(16)).toNumber()] = (_725_v7).IsSubsetOf(_725_v7);
      _nw115[(new BigNumber(17)).toNumber()] = !(_727_v9).contains(_726_v8);
      _nw115[(new BigNumber(18)).toNumber()] = _719_v1;
      _nw115[(new BigNumber(19)).toNumber()] = _719_v1;
      _nw115[(new BigNumber(20)).toNumber()] = (_719_v1) || (_719_v1);
      _nw115[(new BigNumber(21)).toNumber()] = !((_719_v1) && (_719_v1));
      _nw115[(new BigNumber(22)).toNumber()] = (_728_v10).IsDisjointFrom(_728_v10);
      _nw115[(new BigNumber(23)).toNumber()] = (_729_v11).IsDisjointFrom(_dafny.MultiSet.fromElements(_719_v1, _719_v1));
      _nw115[(new BigNumber(24)).toNumber()] = false;
      _nw115[(new BigNumber(25)).toNumber()] = _719_v1;
      _nw115[(new BigNumber(26)).toNumber()] = (((_731_v13).contains(_dafny.MultiSet.fromElements((_717_v0)[_module.__default.safeIndex(new BigNumber(993), new BigNumber((_717_v0).length))], (_717_v0)[_module.__default.safeIndex(new BigNumber(993), new BigNumber((_717_v0).length))], p0, (_this).f7))) ? ((_731_v13).get(_dafny.MultiSet.fromElements((_717_v0)[_module.__default.safeIndex(new BigNumber(993), new BigNumber((_717_v0).length))], (_717_v0)[_module.__default.safeIndex(new BigNumber(993), new BigNumber((_717_v0).length))], p0, (_this).f7))) : (!(_719_v1)));
      _732_v14 = _nw115;
      let _index114 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_732_v14).length));
      (_732_v14)[_index114] = ((_719_v1) ? (_719_v1) : (!(_719_v1)));
      let _733_v15;
      _733_v15 = _dafny.MultiSet.fromElements((_this).f8);
      let _index115 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_717_v0).length));
      let _index116 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_732_v14).length));
      let _rhs121 = (_this).fm7(_module.__default.safeModuloInt(new BigNumber(((_733_v15).update((_717_v0)[_module.__default.safeIndex(new BigNumber(993), new BigNumber((_717_v0).length))], _module.__default.abs((_this).f7))).cardinality()), p0), globalState);
      let _rhs122 = (_720_v2)[_module.__default.safeIndex((_this).f7, new BigNumber((_720_v2).length))];
      let _lhs87 = _717_v0;
      let _lhs88 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_717_v0).length));
      let _lhs89 = _732_v14;
      let _lhs90 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_732_v14).length));
      _lhs87[_lhs88] = _rhs121;
      _lhs89[_lhs90] = _rhs122;
      let _734_v16;
      _734_v16 = _dafny.Seq.of(_717_v0, _717_v0);
      let _index117 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_717_v0).length));
      (_717_v0)[_index117] = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_734_v16, _734_v16), _module.__default.safeIndex((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_732_v14)[_module.__default.safeIndex(new BigNumber(888), new BigNumber((_732_v14).length))],(_732_v14)[_module.__default.safeIndex(new BigNumber(888), new BigNumber((_732_v14).length))])).length)).plus((_this).f8), new BigNumber((_dafny.Seq.Concat(_734_v16, _734_v16)).length)), _717_v0)).length);
      if (_719_v1) {
        let _735_v17;
        _735_v17 = _dafny.Map.Empty.slice().updateUnsafe(_722_v4,_dafny.Seq.update(_module.__default.fm10((_this).f9, (_732_v14)[_module.__default.safeIndex(new BigNumber(888), new BigNumber((_732_v14).length))], (_dafny.ZERO).minus(p0), (_this).f7, globalState), _module.__default.safeIndex((_this).f7, new BigNumber((_module.__default.fm10((_this).f9, (_732_v14)[_module.__default.safeIndex(new BigNumber(888), new BigNumber((_732_v14).length))], (_dafny.ZERO).minus(p0), (_this).f7, globalState)).length)), _719_v1));
        let _736_v18;
        _736_v18 = _dafny.Seq.of(_724_v6, _724_v6);
        _735_v17 = (_735_v17).update(_722_v4, _736_v18);
        _719_v1 = (((_726_v8).contains((_732_v14)[_module.__default.safeIndex(new BigNumber(888), new BigNumber((_732_v14).length))])) ? ((_726_v8).get((_732_v14)[_module.__default.safeIndex(new BigNumber(888), new BigNumber((_732_v14).length))])) : (!((_732_v14)[_module.__default.safeIndex(new BigNumber(888), new BigNumber((_732_v14).length))])));
        let _index118 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_732_v14).length));
        (_732_v14)[_index118] = _719_v1;
        let _737_v19;
        _737_v19 = _dafny.Map.Empty.slice().updateUnsafe(_719_v1,_732_v14);
        let _738_v20;
        let _nw116 = Array((new BigNumber(15)).toNumber());
        _nw116[(_dafny.ZERO).toNumber()] = _732_v14;
        _nw116[(_dafny.ONE).toNumber()] = _732_v14;
        _nw116[(new BigNumber(2)).toNumber()] = _732_v14;
        _nw116[(new BigNumber(3)).toNumber()] = _732_v14;
        _nw116[(new BigNumber(4)).toNumber()] = (((_737_v19).contains(_719_v1)) ? ((_737_v19).get(_719_v1)) : (_732_v14));
        _nw116[(new BigNumber(5)).toNumber()] = _732_v14;
        _nw116[(new BigNumber(6)).toNumber()] = _732_v14;
        _nw116[(new BigNumber(7)).toNumber()] = _732_v14;
        _nw116[(new BigNumber(8)).toNumber()] = _732_v14;
        _nw116[(new BigNumber(9)).toNumber()] = _732_v14;
        _nw116[(new BigNumber(10)).toNumber()] = _732_v14;
        _nw116[(new BigNumber(11)).toNumber()] = _732_v14;
        _nw116[(new BigNumber(12)).toNumber()] = (((_737_v19).contains(true)) ? ((_737_v19).get(true)) : (_732_v14));
        _nw116[(new BigNumber(13)).toNumber()] = _732_v14;
        _nw116[(new BigNumber(14)).toNumber()] = _732_v14;
        _738_v20 = _nw116;
        let _index119 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_738_v20).length));
        (_738_v20)[_index119] = _732_v14;
        let _index120 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_738_v20).length));
        let _nw117 = Array((new BigNumber(23)).toNumber()).fill(false);
        (_738_v20)[_index120] = _nw117;
        _732_v14 = (_738_v20)[_module.__default.safeIndex(new BigNumber(888), new BigNumber((_738_v20).length))];
      } else {
        let _739_v22;
        _739_v22 = _dafny.Set.fromElements((_this).f9);
        let _index121 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_732_v14).length));
        let _rhs123 = (_this).f7;
        let _rhs124 = (_739_v22).IsProperSubsetOf(function () {
          let _coll35 = new _dafny.Set();
          for (const _compr_35 of (_dafny.Seq.of(_dafny.Seq.update((_this).f9, _module.__default.safeIndex((_717_v0)[_module.__default.safeIndex(new BigNumber(993), new BigNumber((_717_v0).length))], new BigNumber(((_this).f9).length)), new _dafny.CodePoint('f'.codePointAt(0))), (_this).f9)).Elements) {
            let _740_v21 = _compr_35;
            if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.update((_this).f9, _module.__default.safeIndex((_717_v0)[_module.__default.safeIndex(new BigNumber(993), new BigNumber((_717_v0).length))], new BigNumber(((_this).f9).length)), new _dafny.CodePoint('f'.codePointAt(0))), (_this).f9), _740_v21)) {
              _coll35.add(_740_v21);
            }
          }
          return _coll35;
        }());
        let _lhs91 = _732_v14;
        let _lhs92 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_732_v14).length));
        r0 = _rhs123;
        _lhs91[_lhs92] = _rhs124;
        let _741_v23;
        let _nw118 = new _module.C3();
        _nw118.__ctor(_dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("xs"), (_this).f9), (_this).f7, (_717_v0)[_module.__default.safeIndex(new BigNumber(993), new BigNumber((_717_v0).length))]);
        _741_v23 = _nw118;
        let _742_v24;
        let _nw119 = new _module.C1();
        _nw119.__ctor(((_this).f7).plus((_this).f7), ((_dafny.ZERO).minus(new BigNumber((_730_v12).length))).minus((_this).f7));
        _742_v24 = _nw119;
        if (_719_v1) {
          _733_v15 = _dafny.MultiSet.fromElements(new BigNumber(-599), (_this).f8, (_this).fm7((_717_v0)[_module.__default.safeIndex(new BigNumber(993), new BigNumber((_717_v0).length))], globalState));
          let _743_v25;
          let _nw120 = new _module.C0();
          _nw120.__ctor(_717_v0, new BigNumber(991), (_this).f10);
          _743_v25 = _nw120;
          let _744_v26;
          _744_v26 = _dafny.Seq.of(_743_v25);
          let _745_v27;
          _745_v27 = _dafny.Set.fromElements(_744_v26);
          (globalState).f3 = (_module.__default.safeModuloInt(_module.__default.fm16(_719_v1, globalState), (_717_v0)[_module.__default.safeIndex(new BigNumber(993), new BigNumber((_717_v0).length))])).isLessThanOrEqualTo(_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_745_v27).length)), p0));
          let _index122 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_717_v0).length));
          (_717_v0)[_index122] = _module.__default.safeModuloInt((_717_v0)[_module.__default.safeIndex(new BigNumber(993), new BigNumber((_717_v0).length))], new BigNumber(965));
          let _index123 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_717_v0).length));
          (_717_v0)[_index123] = new BigNumber(602);
          let _746_v28;
          let _nw121 = new _module.C4();
          _nw121.__ctor((_741_v23).f12, (_this).f10, p0);
          _746_v28 = _nw121;
        } else {
          _730_v12 = _dafny.Seq.Concat(_730_v12, _730_v12);
          let _747_v29;
          let _nw122 = new _module.C1();
          _nw122.__ctor((_this).f8, (_dafny.ZERO).minus((_this).fm7((_this).f8, globalState)));
          _747_v29 = _nw122;
          let _748_v30;
          _748_v30 = _dafny.Map.Empty.slice().updateUnsafe((_741_v23).f12,new BigNumber((_720_v2).length));
          (globalState).f4 = !((_748_v30).update((_741_v23).f12, (_this).f10)).equals(_748_v30);
          (globalState).f4 = (_732_v14)[_module.__default.safeIndex(new BigNumber(888), new BigNumber((_732_v14).length))];
          let _index124 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_717_v0).length));
          (_717_v0)[_index124] = (_this).f10;
        }
        let _749_v31;
        _749_v31 = _dafny.Map.Empty.slice().updateUnsafe((_733_v15).Union(_733_v15),_dafny.Seq.UnicodeFromString("fvhdv"));
        _749_v31 = (_749_v31).update(_dafny.MultiSet.FromArray(_730_v12), _dafny.Seq.Concat((_this).f9, (_this).f9));
      }
      let _750_v32;
      let _nw123 = new _module.C4();
      _nw123.__ctor((_732_v14)[_module.__default.safeIndex(new BigNumber(888), new BigNumber((_732_v14).length))], (_717_v0)[_module.__default.safeIndex(new BigNumber(993), new BigNumber((_717_v0).length))], (_this).f10);
      _750_v32 = _nw123;
      let _751_v33;
      _751_v33 = _dafny.Seq.of(_750_v32, _750_v32);
      _751_v33 = _751_v33;
      let _hi7 = (_this).f7;
      for (let _752_i1 = p0; _752_i1.isLessThan(_hi7); _752_i1 = _752_i1.plus(_dafny.ONE)) {
        let _753_v34;
        _753_v34 = _dafny.Seq.UnicodeFromString("tjdnlna");
        let _754_v35;
        let _nw124 = Array((new BigNumber(14)).toNumber()).fill([]);
        _754_v35 = _nw124;
        let _index125 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_754_v35).length));
        (_754_v35)[_index125] = _732_v14;
        let _755_v36;
        let _nw125 = Array((new BigNumber(2)).toNumber());
        _nw125[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('f'.codePointAt(0));
        _nw125[(_dafny.ONE).toNumber()] = _722_v4;
        _755_v36 = _nw125;
        let _index126 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_755_v36).length));
        (_755_v36)[_index126] = _722_v4;
        let _756_v37;
        _756_v37 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(p0, new BigNumber((_733_v15).cardinality())),(_732_v14)[_module.__default.safeIndex(new BigNumber(888), new BigNumber((_732_v14).length))]);
        let _index127 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_717_v0).length));
        let _index128 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_754_v35).length));
        let _index129 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_755_v36).length));
        let _rhs125 = _dafny.Seq.UnicodeFromString("olimw");
        let _rhs126 = new BigNumber(((_756_v37).Merge(_756_v37)).length);
        let _rhs127 = ((_this).f8).isLessThan(_module.__default.safeModuloInt((_this).f8, (_this).f8));
        let _rhs128 = _732_v14;
        let _rhs129 = _722_v4;
        let _lhs93 = _717_v0;
        let _lhs94 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_717_v0).length));
        let _lhs95 = globalState;
        let _lhs96 = _754_v35;
        let _lhs97 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_754_v35).length));
        let _lhs98 = _755_v36;
        let _lhs99 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_755_v36).length));
        _753_v34 = _rhs125;
        _lhs93[_lhs94] = _rhs126;
        _lhs95.f3 = _rhs127;
        _lhs96[_lhs97] = _rhs128;
        _lhs98[_lhs99] = _rhs129;
        let _757_v38;
        _757_v38 = _module.D3.create_DC7((_750_v32).f11);
        let _source17 = _757_v38;
        if (_source17.is_DC7) {
          let _758___mcc_h0 = (_source17).cf10;
          let _759_cf10 = _758___mcc_h0;
          let _760_v39;
          _760_v39 = _dafny.Set.fromElements(_723_v5, _723_v5, _723_v5, _723_v5, _723_v5);
          let _761_v41;
          let _nw126 = Array((new BigNumber(15)).toNumber());
          _nw126[(_dafny.ZERO).toNumber()] = _module.__default.fm31((_this).f8, _dafny.MultiSet.FromArray(_module.__default.fm30(_759_cf10, globalState)), globalState);
          _nw126[(_dafny.ONE).toNumber()] = _760_v39;
          _nw126[(new BigNumber(2)).toNumber()] = function () {
            let _coll36 = new _dafny.Set();
            for (const _compr_36 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(386)), ((_762_v5, _763_i1, _764_p0) => function (_765_i2) {
              return (_762_v5).update(new BigNumber((_dafny.Seq.of(_763_i1)).length), _764_p0);
            })(_723_v5, _752_i1, p0))).Elements) {
              let _766_v40 = _compr_36;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(386)), ((_767_v5, _768_i1, _769_p0) => function (_765_i2) {
                return (_767_v5).update(new BigNumber((_dafny.Seq.of(_768_i1)).length), _769_p0);
              })(_723_v5, _752_i1, p0)), _766_v40)) {
                _coll36.add(_766_v40);
              }
            }
            return _coll36;
          }();
          _nw126[(new BigNumber(3)).toNumber()] = _760_v39;
          _nw126[(new BigNumber(4)).toNumber()] = _760_v39;
          _nw126[(new BigNumber(5)).toNumber()] = _760_v39;
          _nw126[(new BigNumber(6)).toNumber()] = _760_v39;
          _nw126[(new BigNumber(7)).toNumber()] = (_760_v39).Intersect(_760_v39);
          _nw126[(new BigNumber(8)).toNumber()] = _760_v39;
          _nw126[(new BigNumber(9)).toNumber()] = (_module.__default.fm31((_717_v0)[_module.__default.safeIndex(new BigNumber(993), new BigNumber((_717_v0).length))], _729_v11, globalState)).Difference(_760_v39);
          _nw126[(new BigNumber(10)).toNumber()] = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_this).f8,(_dafny.ZERO).minus((_this).f10)));
          _nw126[(new BigNumber(11)).toNumber()] = _760_v39;
          _nw126[(new BigNumber(12)).toNumber()] = _760_v39;
          _nw126[(new BigNumber(13)).toNumber()] = _760_v39;
          _nw126[(new BigNumber(14)).toNumber()] = _module.__default.fm31((_this).f7, _dafny.MultiSet.FromArray(_720_v2), globalState);
          _761_v41 = _nw126;
          let _index130 = _module.__default.safeIndex(new BigNumber(196), new BigNumber((_761_v41).length));
          (_761_v41)[_index130] = _760_v39;
          let _770_v42;
          _770_v42 = _dafny.Seq.of((_760_v39).Intersect(_760_v39));
          let _index131 = _module.__default.safeIndex(new BigNumber(196), new BigNumber((_761_v41).length));
          let _index132 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_717_v0).length));
          let _index133 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_755_v36).length));
          let _rhs130 = (_this).f9;
          let _rhs131 = ((_this).f8).isLessThan((_this).f10);
          let _rhs132 = (_770_v42)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f10), new BigNumber((_770_v42).length))];
          let _rhs133 = ((_this).f7).plus(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.UnicodeFromString("cniy")).length), new BigNumber((_dafny.Seq.update((_this).f9, _module.__default.safeIndex((_dafny.ZERO).minus((_this).f8), new BigNumber(((_this).f9).length)), (_755_v36)[_module.__default.safeIndex(new BigNumber(652), new BigNumber((_755_v36).length))])).length)));
          let _rhs134 = _722_v4;
          let _lhs100 = _761_v41;
          let _lhs101 = _module.__default.safeIndex(new BigNumber(196), new BigNumber((_761_v41).length));
          let _lhs102 = _717_v0;
          let _lhs103 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_717_v0).length));
          let _lhs104 = _755_v36;
          let _lhs105 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_755_v36).length));
          _753_v34 = _rhs130;
          _719_v1 = _rhs131;
          _lhs100[_lhs101] = _rhs132;
          _lhs102[_lhs103] = _rhs133;
          _lhs104[_lhs105] = _rhs134;
          _717_v0 = _717_v0;
          let _index134 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_732_v14).length));
          let _rhs135 = true;
          let _rhs136 = _717_v0;
          let _rhs137 = _dafny.Seq.of((_730_v12)[_module.__default.safeIndex(new BigNumber((_733_v15).cardinality()), new BigNumber((_730_v12).length))]);
          let _rhs138 = ((_750_v32).f11) === ((_732_v14)[_module.__default.safeIndex(new BigNumber(888), new BigNumber((_732_v14).length))]);
          let _rhs139 = _727_v9;
          let _lhs106 = _732_v14;
          let _lhs107 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_732_v14).length));
          _lhs106[_lhs107] = _rhs135;
          _717_v0 = _rhs136;
          _730_v12 = _rhs137;
          _759_cf10 = _rhs138;
          _727_v9 = _rhs139;
          _721_v3 = (_721_v3).update((_this).f9, (_732_v14)[_module.__default.safeIndex(new BigNumber(888), new BigNumber((_732_v14).length))]);
        } else {
          let _771___mcc_h1 = (_source17).cf9;
          let _772_cf9 = _771___mcc_h1;
          (globalState).f4 = _719_v1;
          let _773_v43;
          _773_v43 = _dafny.Map.Empty.slice().updateUnsafe(_725_v7,_dafny.Seq.Concat((_this).f9, _753_v34));
          _753_v34 = (((_773_v43).contains((_772_cf9).Union(_772_cf9))) ? ((_773_v43).get((_772_cf9).Union(_772_cf9))) : (_753_v34));
          let _index135 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_717_v0).length));
          (_717_v0)[_index135] = _752_i1;
          (globalState).f4 = (_750_v32).f11;
        }
        let _774_v44;
        _774_v44 = _module.D1.create_DC1(new _dafny.CodePoint('u'.codePointAt(0)));
        let _775_v45;
        _775_v45 = _module.D1.create_DC3(_774_v44);
        let _776_v46;
        let _777_v47;
        let _778_v48;
        let _779_v49;
        let _out6;
        let _out7;
        let _out8;
        let _out9;
        let _outcollector3 = (_750_v32).m5((_750_v32).f11, new BigNumber((_730_v12).length), new BigNumber(677), new BigNumber((_dafny.Set.fromElements(_775_v45)).length), globalState);
        _out6 = _outcollector3[0];
        _out7 = _outcollector3[1];
        _out8 = _outcollector3[2];
        _out9 = _outcollector3[3];
        _776_v46 = _out6;
        _777_v47 = _out7;
        _778_v48 = _out8;
        _779_v49 = _out9;
        r0 = _module.__default.safeDivisionInt(new BigNumber(((_726_v8).Merge(_726_v8)).length), (_717_v0)[_module.__default.safeIndex(new BigNumber(993), new BigNumber((_717_v0).length))]);
      }
      r0 = (_this).f8;
      return r0;
    }
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f10() {
      let _this = this;
      return _this._f10;
    };
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this._f7 = _dafny.ZERO;
      this._f8 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T2, _module.T1, _module.T0];
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    __ctor(f7, f8) {
      let _this = this;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      return;
    }
    fm2(p0, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("lysobld"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(732)), function (_780_i0) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      })), _dafny.Seq.UnicodeFromString("exbusrxu"));
    };
    fm0(p0, globalState) {
      let _this = this;
      return (_dafny.Set.fromElements(_dafny.Set.fromElements(false, !(!(true)), !(false)))).Intersect((_dafny.Set.fromElements(_dafny.Set.fromElements(true, false))).Difference(_dafny.Set.fromElements(_dafny.Set.fromElements(false), _dafny.Set.fromElements(false, true, false, true, true))));
    };
    fm1(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (new BigNumber(970)).isLessThanOrEqualTo((_this).f7);
    };
    fm3(p0, p1, p2, p3, globalState) {
      let _this = this;
      if ((!(false))) {
        return new BigNumber(-377);
      } else {
        return new BigNumber((_dafny.Seq.UnicodeFromString("sv")).length);
      }
    };
    fm4(p0, p1, globalState) {
      let _this = this;
      return (_this).f7;
    };
    m1(p0, globalState) {
      let _this = this;
      let _781_v0;
      _781_v0 = new BigNumber(-479);
      let _782_v1;
      _782_v1 = _dafny.Seq.UnicodeFromString("hndygtt");
      let _783_v2;
      _783_v2 = true;
      let _784_v3;
      _784_v3 = _dafny.MultiSet.fromElements(_783_v2);
      let _rhs140 = _783_v2;
      let _rhs141 = ((_783_v2) ? (((_this).f8).multipliedBy((((_784_v3).contains(_783_v2)) ? ((_784_v3).get(_783_v2)) : ((_this).f8)))) : (((_this).f7).plus((_this).f8)));
      let _rhs142 = _dafny.Seq.update(_782_v1, _module.__default.safeIndex(_module.__default.safeDivisionInt((_this).f7, _781_v0), new BigNumber((_782_v1).length)), _module.__default.fm5(globalState));
      let _lhs108 = globalState;
      _lhs108.f3 = _rhs140;
      _781_v0 = _rhs141;
      _782_v1 = _rhs142;
      let _785_v4;
      let _nw127 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
      _785_v4 = _nw127;
      let _786_v5;
      _786_v5 = _dafny.MultiSet.fromElements(new BigNumber((_782_v1).length));
      let _787_v6;
      _787_v6 = _dafny.Map.Empty.slice().updateUnsafe(_783_v2,_786_v5);
      let _788_v7;
      _788_v7 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(((((_787_v6).contains(_783_v2)) ? ((_787_v6).get(_783_v2)) : (_786_v5))).cardinality())), (_this).f7);
      let _789_v8;
      let _nw128 = new _module.C0();
      _nw128.__ctor(_785_v4, (_this).f8, _module.__default.safeModuloInt((_this).f8, (_788_v7)[_module.__default.safeIndex((_this).f7, new BigNumber((_788_v7).length))]));
      _789_v8 = _nw128;
      _789_v8 = _789_v8;
      let _790_v9;
      _790_v9 = _dafny.Seq.of((_786_v5).update(_781_v0, _module.__default.abs(_781_v0)));
      let _791_i0;
      _791_i0 = _dafny.ZERO;
      L5: {
        while (_dafny.Seq.IsProperPrefixOf(_790_v9, _dafny.Seq.of(_786_v5, _786_v5))) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_791_i0)) {
              break L5;
            }
            _791_i0 = (_791_i0).plus(_dafny.ONE);
            if (((_783_v2) ? (_dafny.areEqual(p0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(938)), function (_792_i1) {
              return new _dafny.CodePoint('e'.codePointAt(0));
            }))) : (((_783_v2) ? (_783_v2) : (_783_v2))))) {
              let _793_v10;
              _793_v10 = new _dafny.CodePoint('r'.codePointAt(0));
              _782_v1 = _dafny.Seq.update(_dafny.Seq.Concat(_782_v1, _782_v1), _module.__default.safeIndex(_module.__default.safeDivisionInt((_789_v8).f7, (_this).f7), new BigNumber((_dafny.Seq.Concat(_782_v1, _782_v1)).length)), _793_v10);
              let _794_v11;
              let _init19 = ((_795_p0, _796_v1) => function (_797_i2) {
                return _dafny.areEqual(_795_p0, _796_v1);
              })(p0, _782_v1);
              let _nw129 = Array((new BigNumber(26)).toNumber());
              for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw129.length); _i0_19++) {
                _nw129[_i0_19] = _init19(new BigNumber(_i0_19));
              }
              _794_v11 = _nw129;
              let _798_v12;
              _798_v12 = _module.D5.create_DC13(true, _793_v10);
              let _799_v13;
              _799_v13 = _module.D6.create_DC15(_794_v11);
              let _rhs143 = (_799_v13).dtor_cf23;
              let _rhs144 = _module.D5.create_DC13(_783_v2, _793_v10);
              _794_v11 = _rhs143;
              _798_v12 = _rhs144;
              let _800_v14;
              let _nw130 = Array((new BigNumber(9)).toNumber()).fill(_module.D2.Default());
              _800_v14 = _nw130;
              let _801_v15;
              _801_v15 = _dafny.Map.Empty.slice().updateUnsafe(_783_v2,_800_v14);
              _800_v14 = (((_801_v15).contains(_783_v2)) ? ((_801_v15).get(_783_v2)) : (_800_v14));
              let _802_v16;
              _802_v16 = _dafny.Set.fromElements(_793_v10, _793_v10);
              _802_v16 = (_802_v16).Union(_802_v16);
              let _803_v17;
              _803_v17 = _dafny.Set.fromElements(_module.__default.fm8(_783_v2, _788_v7, _783_v2, globalState), _dafny.Seq.Concat(p0, p0));
              let _804_v18;
              _804_v18 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-863)), ((_805_v10) => function (_806_i3) {
                return _805_v10;
              })(_793_v10)),_dafny.Seq.of(_783_v2));
              _803_v17 = ((_803_v17).Difference(_dafny.Set.fromElements(_782_v1))).Union(function () {
                let _coll37 = new _dafny.Set();
                for (const _compr_37 of (_804_v18).Keys.Elements) {
                  let _807_v19 = _compr_37;
                  if ((_804_v18).contains(_807_v19)) {
                    _coll37.add(_807_v19);
                  }
                }
                return _coll37;
              }());
            } else {
              let _index136 = _module.__default.safeIndex(new BigNumber(65), new BigNumber((_785_v4).length));
              (_785_v4)[_index136] = (((_784_v3).contains(!(_783_v2))) ? ((_784_v3).get(!(_783_v2))) : ((_this).f7));
              let _index137 = _module.__default.safeIndex(new BigNumber(65), new BigNumber((_785_v4).length));
              (_785_v4)[_index137] = _module.__default.safeDivisionInt(_module.__default.fm20(_783_v2, _783_v2, _783_v2, globalState), (_789_v8).f8);
              let _808_v20;
              _808_v20 = _module.D7.create_DC17(p0);
              let _809_v21;
              _809_v21 = _dafny.Seq.of(_782_v1, _module.__default.fm8(_783_v2, _788_v7, _783_v2, globalState));
              let _810_v22;
              _810_v22 = _dafny.Set.fromElements(_782_v1, _dafny.Seq.UnicodeFromString("j"));
              let _811_v23;
              let _nw131 = Array((new BigNumber(21)).toNumber());
              _nw131[(_dafny.ZERO).toNumber()] = _782_v1;
              _nw131[(_dafny.ONE).toNumber()] = _782_v1;
              _nw131[(new BigNumber(2)).toNumber()] = _782_v1;
              _nw131[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_782_v1, _782_v1);
              _nw131[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(362)), function (_812_i4) {
                return new _dafny.CodePoint('b'.codePointAt(0));
              });
              _nw131[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_782_v1, p0);
              _nw131[(new BigNumber(6)).toNumber()] = (_808_v20).dtor_cf26;
              _nw131[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(p0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(10)), function (_813_i5) {
                return new _dafny.CodePoint('l'.codePointAt(0));
              }));
              _nw131[(new BigNumber(8)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(805)), function (_814_i6) {
                return new _dafny.CodePoint('u'.codePointAt(0));
              });
              _nw131[(new BigNumber(9)).toNumber()] = _782_v1;
              _nw131[(new BigNumber(10)).toNumber()] = p0;
              _nw131[(new BigNumber(11)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(519)), function (_815_i7) {
                return new _dafny.CodePoint('o'.codePointAt(0));
              });
              _nw131[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat((_808_v20).dtor_cf26, _dafny.Seq.UnicodeFromString("klcneosa"));
              _nw131[(new BigNumber(13)).toNumber()] = _dafny.Seq.UnicodeFromString("vbsipip");
              _nw131[(new BigNumber(14)).toNumber()] = _dafny.Seq.UnicodeFromString("wudxpkey");
              _nw131[(new BigNumber(15)).toNumber()] = p0;
              _nw131[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_782_v1, (_809_v21)[_module.__default.safeIndex(new BigNumber((_810_v22).length), new BigNumber((_809_v21).length))]);
              _nw131[(new BigNumber(17)).toNumber()] = _782_v1;
              _nw131[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_782_v1, _782_v1);
              _nw131[(new BigNumber(19)).toNumber()] = _dafny.Seq.update(p0, _module.__default.safeIndex(new BigNumber((_786_v5).cardinality()), new BigNumber((p0).length)), _module.__default.fm5(globalState));
              _nw131[(new BigNumber(20)).toNumber()] = _782_v1;
              _811_v23 = _nw131;
              let _index138 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_811_v23).length));
              (_811_v23)[_index138] = _782_v1;
              let _816_v24;
              _816_v24 = _dafny.Seq.of(_783_v2);
              let _817_v25;
              _817_v25 = _dafny.Map.Empty.slice().updateUnsafe(_781_v0,(_module.D5.create_DC12(_783_v2)).dtor_cf19);
              let _818_v26;
              _818_v26 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_816_v24).length),new BigNumber((_817_v25).length));
              let _819_v27;
              _819_v27 = new _dafny.CodePoint('j'.codePointAt(0));
              let _index139 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_811_v23).length));
              (_811_v23)[_index139] = _dafny.Seq.Concat(_dafny.Seq.update(_782_v1, _module.__default.safeIndex(new BigNumber((_818_v26).length), new BigNumber((_782_v1).length)), _819_v27), _dafny.Seq.UnicodeFromString("vhe"));
              let _820_v28;
              let _nw132 = Array((new BigNumber(12)).toNumber()).fill(false);
              _820_v28 = _nw132;
              let _index140 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_820_v28).length));
              (_820_v28)[_index140] = _783_v2;
              let _index141 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_820_v28).length));
              (_820_v28)[_index141] = ((!(!(_783_v2) || (_783_v2))) ? (_783_v2) : (_783_v2));
              let _821_v29;
              let _init20 = ((_822_v24) => function (_823_i8) {
                return _dafny.Seq.of(_822_v24);
              })(_816_v24);
              let _nw133 = Array((new BigNumber(9)).toNumber());
              for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw133.length); _i0_20++) {
                _nw133[_i0_20] = _init20(new BigNumber(_i0_20));
              }
              _821_v29 = _nw133;
              let _824_v30;
              _824_v30 = _dafny.Seq.of(_816_v24, _816_v24);
              let _index142 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_821_v29).length));
              (_821_v29)[_index142] = _dafny.Seq.Concat(_824_v30, _824_v30);
              let _825_v31;
              _825_v31 = _module.D5.create_DC13((_820_v28)[_module.__default.safeIndex(new BigNumber(300), new BigNumber((_820_v28).length))], _819_v27);
              let _826_v32;
              _826_v32 = _dafny.Map.Empty.slice().updateUnsafe(_825_v31,_819_v27);
              let _827_v33;
              _827_v33 = _module.D7.create_DC20(_819_v27, new BigNumber((((true) ? (_826_v32) : (_826_v32))).length));
              let _index143 = _module.__default.safeIndex(new BigNumber(65), new BigNumber((_785_v4).length));
              let _index144 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_821_v29).length));
              let _rhs145 = (_789_v8).f7;
              let _rhs146 = _824_v30;
              let _rhs147 = _783_v2;
              let _rhs148 = _module.D7.create_DC20(_819_v27, (_dafny.ZERO).minus((_dafny.ZERO).minus((_789_v8).f8)));
              let _rhs149 = !((_820_v28)[_module.__default.safeIndex(new BigNumber(300), new BigNumber((_820_v28).length))]);
              let _lhs109 = _785_v4;
              let _lhs110 = _module.__default.safeIndex(new BigNumber(65), new BigNumber((_785_v4).length));
              let _lhs111 = _821_v29;
              let _lhs112 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_821_v29).length));
              let _lhs113 = globalState;
              let _lhs114 = globalState;
              _lhs109[_lhs110] = _rhs145;
              _lhs111[_lhs112] = _rhs146;
              _lhs113.f3 = _rhs147;
              _827_v33 = _rhs148;
              _lhs114.f4 = _rhs149;
              _781_v0 = (_789_v8).f7;
            }
            let _828_v34;
            _828_v34 = new _dafny.CodePoint('k'.codePointAt(0));
            let _829_v35;
            _829_v35 = _module.D5.create_DC13(_783_v2, _828_v34);
            let _source18 = _829_v35;
            if (_source18.is_DC12) {
              let _830___mcc_h0 = (_source18).cf19;
              let _831_cf19 = _830___mcc_h0;
              let _index145 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_785_v4).length));
              (_785_v4)[_index145] = (_dafny.ZERO).minus((_789_v8).f7);
              let _index146 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_785_v4).length));
              (_785_v4)[_index146] = (_this).f8;
              _781_v0 = (_789_v8).f7;
              let _832_v36;
              _832_v36 = _dafny.Map.Empty.slice().updateUnsafe(_831_cf19,(_785_v4)[_module.__default.safeIndex(new BigNumber(170), new BigNumber((_785_v4).length))]);
              let _index147 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_785_v4).length));
              (_785_v4)[_index147] = _module.__default.safeDivisionInt(new BigNumber((_832_v36).length), (_785_v4)[_module.__default.safeIndex(new BigNumber(170), new BigNumber((_785_v4).length))]);
              let _833_v37;
              let _init21 = ((_834_v2) => function (_835_i9) {
                return _834_v2;
              })(_783_v2);
              let _nw134 = Array((new BigNumber(19)).toNumber());
              for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw134.length); _i0_21++) {
                _nw134[_i0_21] = _init21(new BigNumber(_i0_21));
              }
              _833_v37 = _nw134;
              let _836_v38;
              _836_v38 = _dafny.Seq.of(_833_v37);
              let _837_v39;
              _837_v39 = _module.D4.create_DC9((_this).f8, _836_v38, (_dafny.ZERO).minus((_this).f8), (_dafny.ZERO).minus((_789_v8).f8));
              let _pat_let_tv12 = _781_v0;
              let _pat_let_tv13 = _781_v0;
              let _838_v40;
              _838_v40 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let9_0) {
                return function (_839_dt__update__tmp_h0) {
                  return function (_pat_let10_0) {
                    return function (_840_dt__update_hcf15_h0) {
                      return function (_pat_let11_0) {
                        return function (_841_dt__update_hcf12_h0) {
                          return _module.D4.create_DC9(_841_dt__update_hcf12_h0, (_839_dt__update__tmp_h0).dtor_cf13, (_839_dt__update__tmp_h0).dtor_cf14, _840_dt__update_hcf15_h0);
                        }(_pat_let11_0);
                      }(_pat_let_tv13);
                    }(_pat_let10_0);
                  }(_pat_let_tv12);
                }(_pat_let9_0);
              }(_837_v39),(_789_v8).f7);
              let _842_v41;
              _842_v41 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,_831_cf19);
              _838_v40 = (_838_v40).update(_837_v39, new BigNumber((_842_v41).length));
            } else if (_source18.is_DC13) {
              let _843___mcc_h1 = (_source18).cf20;
              let _844___mcc_h2 = (_source18).cf21;
              let _845_cf21 = _844___mcc_h2;
              let _846_cf20 = _843___mcc_h1;
              let _847_v42;
              _847_v42 = _dafny.Set.fromElements(!(_846_cf20));
              let _848_v43;
              _848_v43 = _dafny.Set.fromElements(_dafny.Set.fromElements(!(_783_v2)), _847_v42, (_847_v42).Difference(_847_v42));
              _848_v43 = (_dafny.Set.fromElements(_847_v42, _847_v42, _847_v42, _dafny.Set.fromElements(_846_cf20))).Union(_848_v43);
              let _849_v44;
              let _nw135 = Array((new BigNumber(15)).toNumber());
              _nw135[(_dafny.ZERO).toNumber()] = _846_cf20;
              _nw135[(_dafny.ONE).toNumber()] = _846_cf20;
              _nw135[(new BigNumber(2)).toNumber()] = _783_v2;
              _nw135[(new BigNumber(3)).toNumber()] = _783_v2;
              _nw135[(new BigNumber(4)).toNumber()] = _846_cf20;
              _nw135[(new BigNumber(5)).toNumber()] = _846_cf20;
              _nw135[(new BigNumber(6)).toNumber()] = true;
              _nw135[(new BigNumber(7)).toNumber()] = true;
              _nw135[(new BigNumber(8)).toNumber()] = _846_cf20;
              _nw135[(new BigNumber(9)).toNumber()] = _783_v2;
              _nw135[(new BigNumber(10)).toNumber()] = _846_cf20;
              _nw135[(new BigNumber(11)).toNumber()] = true;
              _nw135[(new BigNumber(12)).toNumber()] = _846_cf20;
              _nw135[(new BigNumber(13)).toNumber()] = true;
              _nw135[(new BigNumber(14)).toNumber()] = _846_cf20;
              _849_v44 = _nw135;
              let _850_v45;
              _850_v45 = _dafny.MultiSet.fromElements(_849_v44);
              (globalState).f3 = (_850_v45).IsSubsetOf(((_846_cf20) ? (_850_v45) : (_850_v45)));
              let _index148 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_785_v4).length));
              (_785_v4)[_index148] = (_this).f8;
              let _index149 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_785_v4).length));
              (_785_v4)[_index149] = (_dafny.ZERO).minus(((_dafny.ZERO).minus(new BigNumber((_784_v3).cardinality()))).multipliedBy((_789_v8).f8));
              let _851_v46;
              let _nw136 = Array((new BigNumber(17)).toNumber());
              _nw136[(_dafny.ZERO).toNumber()] = _849_v44;
              _nw136[(_dafny.ONE).toNumber()] = _849_v44;
              _nw136[(new BigNumber(2)).toNumber()] = _849_v44;
              _nw136[(new BigNumber(3)).toNumber()] = _849_v44;
              _nw136[(new BigNumber(4)).toNumber()] = _849_v44;
              _nw136[(new BigNumber(5)).toNumber()] = _849_v44;
              _nw136[(new BigNumber(6)).toNumber()] = _849_v44;
              _nw136[(new BigNumber(7)).toNumber()] = _849_v44;
              _nw136[(new BigNumber(8)).toNumber()] = _849_v44;
              _nw136[(new BigNumber(9)).toNumber()] = _849_v44;
              _nw136[(new BigNumber(10)).toNumber()] = _849_v44;
              _nw136[(new BigNumber(11)).toNumber()] = _849_v44;
              _nw136[(new BigNumber(12)).toNumber()] = _849_v44;
              _nw136[(new BigNumber(13)).toNumber()] = _849_v44;
              _nw136[(new BigNumber(14)).toNumber()] = _849_v44;
              _nw136[(new BigNumber(15)).toNumber()] = _849_v44;
              _nw136[(new BigNumber(16)).toNumber()] = _849_v44;
              _851_v46 = _nw136;
              let _index150 = _module.__default.safeIndex(new BigNumber(149), new BigNumber((_851_v46).length));
              (_851_v46)[_index150] = _849_v44;
              let _index151 = _module.__default.safeIndex(new BigNumber(149), new BigNumber((_851_v46).length));
              (_851_v46)[_index151] = (((_847_v42).IsProperSubsetOf(_dafny.Set.fromElements(_846_cf20, _783_v2, _783_v2, _783_v2, _846_cf20))) ? (_849_v44) : (_849_v44));
            } else if (_source18.is_DC11) {
              let _852___mcc_h3 = (_source18).cf18;
              let _853_cf18 = _852___mcc_h3;
              let _854_v47;
              let _nw137 = new _module.C5();
              _nw137.__ctor(_dafny.Seq.UnicodeFromString("gk"), (_789_v8).f8, (_dafny.ZERO).minus((_this).fm3(_783_v2, _783_v2, _783_v2, _783_v2, globalState)), (_789_v8).f7);
              _854_v47 = _nw137;
              _854_v47 = _854_v47;
              let _855_v48;
              _855_v48 = _dafny.Seq.of(_788_v7);
              _855_v48 = _855_v48;
              let _856_v49;
              let _nw138 = new _module.C3();
              _nw138.__ctor(_783_v2, _module.__default.safeDivisionInt(new BigNumber((p0).length), (_789_v8).f7), (_dafny.ZERO).minus((_789_v8).f8));
              _856_v49 = _nw138;
              let _nw139 = new _module.C2();
              _nw139.__ctor(new BigNumber(922), (_dafny.ZERO).minus((_789_v8).f8));
              _856_v49 = _nw139;
              let _857_v50;
              let _858_v51;
              let _859_v52;
              let _out10;
              let _out11;
              let _out12;
              let _outcollector4 = (_this).m3(((_856_v49).f8).isLessThan((_this).f7), _781_v0, ((_783_v2) ? (_785_v4) : (_785_v4)), (_this).fm4(_783_v2, _783_v2, globalState), globalState);
              _out10 = _outcollector4[0];
              _out11 = _outcollector4[1];
              _out12 = _outcollector4[2];
              _857_v50 = _out10;
              _858_v51 = _out11;
              _859_v52 = _out12;
            } else {
              let _860___mcc_h4 = (_source18).cf22;
              let _861_cf22 = _860___mcc_h4;
              let _862_v53;
              _862_v53 = _module.D1.create_DC2();
              let _863_v54;
              _863_v54 = _dafny.Seq.of(_862_v53);
              _863_v54 = _863_v54;
              let _864_v55;
              _864_v55 = _dafny.Map.Empty.slice().updateUnsafe(_828_v34,(_789_v8).f8);
              _781_v0 = new BigNumber((_864_v55).length);
              _781_v0 = _781_v0;
              _781_v0 = new BigNumber((p0).length);
            }
            let _865_v56;
            _865_v56 = _dafny.Map.Empty.slice().updateUnsafe(_781_v0,_783_v2);
            let _866_v57;
            _866_v57 = _dafny.Map.Empty.slice().updateUnsafe(_865_v56,_785_v4);
            let _867_v58;
            _867_v58 = _dafny.Map.Empty.slice().updateUnsafe(_785_v4,_828_v34);
            let _868_v59;
            _868_v59 = _dafny.Map.Empty.slice().updateUnsafe(_828_v34,_785_v4);
            let _869_v60;
            let _nw140 = Array((new BigNumber(14)).toNumber());
            _nw140[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((((_866_v57).contains(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_788_v7).length),_783_v2))) ? ((_866_v57).get(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_788_v7).length),_783_v2))) : (_785_v4)),_828_v34);
            _nw140[(_dafny.ONE).toNumber()] = _867_v58;
            _nw140[(new BigNumber(2)).toNumber()] = _867_v58;
            _nw140[(new BigNumber(3)).toNumber()] = (_867_v58).Merge(_867_v58);
            _nw140[(new BigNumber(4)).toNumber()] = _867_v58;
            _nw140[(new BigNumber(5)).toNumber()] = _867_v58;
            _nw140[(new BigNumber(6)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_785_v4,_828_v34);
            _nw140[(new BigNumber(7)).toNumber()] = (_867_v58).update((((_868_v59).contains(_828_v34)) ? ((_868_v59).get(_828_v34)) : (_785_v4)), _828_v34);
            _nw140[(new BigNumber(8)).toNumber()] = (_867_v58).Merge(_dafny.Map.Empty.slice().updateUnsafe(_785_v4,_828_v34));
            _nw140[(new BigNumber(9)).toNumber()] = _867_v58;
            _nw140[(new BigNumber(10)).toNumber()] = ((_783_v2) ? (_867_v58) : (_867_v58));
            _nw140[(new BigNumber(11)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_785_v4,_828_v34)).Merge(_867_v58);
            _nw140[(new BigNumber(12)).toNumber()] = _867_v58;
            _nw140[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_785_v4,_828_v34);
            _869_v60 = _nw140;
            let _index152 = _module.__default.safeIndex(new BigNumber(75), new BigNumber((_869_v60).length));
            (_869_v60)[_index152] = _dafny.Map.Empty.slice().updateUnsafe(_785_v4,_828_v34);
            let _index153 = _module.__default.safeIndex(new BigNumber(75), new BigNumber((_869_v60).length));
            (_869_v60)[_index153] = _867_v58;
            let _870_v61;
            _870_v61 = _dafny.Seq.of(_785_v4, _785_v4);
            let _871_v62;
            let _nw141 = Array((new BigNumber(17)).toNumber());
            _nw141[(_dafny.ZERO).toNumber()] = _785_v4;
            _nw141[(_dafny.ONE).toNumber()] = _785_v4;
            _nw141[(new BigNumber(2)).toNumber()] = _785_v4;
            _nw141[(new BigNumber(3)).toNumber()] = _785_v4;
            _nw141[(new BigNumber(4)).toNumber()] = _785_v4;
            _nw141[(new BigNumber(5)).toNumber()] = _785_v4;
            _nw141[(new BigNumber(6)).toNumber()] = _785_v4;
            _nw141[(new BigNumber(7)).toNumber()] = _785_v4;
            _nw141[(new BigNumber(8)).toNumber()] = _785_v4;
            _nw141[(new BigNumber(9)).toNumber()] = _785_v4;
            _nw141[(new BigNumber(10)).toNumber()] = _785_v4;
            _nw141[(new BigNumber(11)).toNumber()] = (_870_v61)[_module.__default.safeIndex((_this).f8, new BigNumber((_870_v61).length))];
            _nw141[(new BigNumber(12)).toNumber()] = _785_v4;
            _nw141[(new BigNumber(13)).toNumber()] = _785_v4;
            _nw141[(new BigNumber(14)).toNumber()] = _785_v4;
            _nw141[(new BigNumber(15)).toNumber()] = _785_v4;
            _nw141[(new BigNumber(16)).toNumber()] = _785_v4;
            _871_v62 = _nw141;
            let _872_v63;
            _872_v63 = _module.D4.create_DC8(_871_v62);
            let _source19 = _872_v63;
            if (_source19.is_DC9) {
              let _873___mcc_h5 = (_source19).cf12;
              let _874___mcc_h6 = (_source19).cf13;
              let _875___mcc_h7 = (_source19).cf14;
              let _876___mcc_h8 = (_source19).cf15;
              let _877_cf15 = _876___mcc_h8;
              let _878_cf14 = _875___mcc_h7;
              let _879_cf13 = _874___mcc_h6;
              let _880_cf12 = _873___mcc_h5;
              let _881_v64;
              _881_v64 = _dafny.Seq.of(_782_v1, _dafny.Seq.of(_828_v34, _828_v34), _782_v1);
              _878_cf14 = _module.__default.safeModuloInt(((_789_v8).f7).multipliedBy(new BigNumber((_784_v3).cardinality())), ((((_784_v3).contains(_783_v2)) ? ((_784_v3).get(_783_v2)) : ((_789_v8).f8))).plus(new BigNumber((_881_v64).length)));
              let _882_v65;
              _882_v65 = _dafny.Map.Empty.slice().updateUnsafe(_783_v2,new BigNumber(647));
              let _883_v66;
              _883_v66 = _dafny.Set.fromElements((_789_v8).f8, _module.__default.safeDivisionInt((_789_v8).f7, new BigNumber(901)), (_dafny.ZERO).minus((_dafny.ZERO).minus((((_882_v65).contains(_783_v2)) ? ((_882_v65).get(_783_v2)) : (new BigNumber((_dafny.Seq.UnicodeFromString("hcam")).length))))), new BigNumber(330));
              _883_v66 = ((_883_v66).Intersect(_dafny.Set.fromElements((_789_v8).f7))).Difference(_883_v66);
              let _884_v67;
              let _nw142 = Array((new BigNumber(20)).toNumber()).fill(false);
              _884_v67 = _nw142;
              _884_v67 = (((_783_v2) === (_783_v2)) ? (_884_v67) : (_884_v67));
              let _index154 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_884_v67).length));
              (_884_v67)[_index154] = _783_v2;
              let _index155 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_884_v67).length));
              (_884_v67)[_index155] = !(_783_v2) || (_783_v2);
            } else if (_source19.is_DC10) {
              let _885___mcc_h9 = (_source19).cf16;
              let _886___mcc_h10 = (_source19).cf17;
              let _887_cf17 = _886___mcc_h10;
              let _888_cf16 = _885___mcc_h9;
              let _889_v68;
              let _nw143 = new _module.C5();
              _nw143.__ctor(_782_v1, new BigNumber(386), new BigNumber((_786_v5).cardinality()), (_789_v8).f7);
              _889_v68 = _nw143;
              let _890_v69;
              _890_v69 = _dafny.Map.Empty.slice().updateUnsafe(_888_cf16,_889_v68);
              let _891_v70;
              _891_v70 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,_890_v69);
              _891_v70 = (_891_v70).update((_789_v8).f7, (_dafny.Map.Empty.slice().updateUnsafe((_this).f7,_889_v68)).update(new BigNumber(358), _889_v68));
              let _892_v71;
              _892_v71 = _dafny.Map.Empty.slice().updateUnsafe(!(_783_v2),(_889_v68).f9);
              let _893_v72;
              _893_v72 = _dafny.Map.Empty.slice().updateUnsafe(true,_783_v2);
              let _894_v73;
              _894_v73 = _module.D7.create_DC20(_828_v34, (_889_v68).f10);
              let _895_v74;
              _895_v74 = _dafny.Seq.of(_894_v73);
              let _896_v75;
              let _nw144 = new _module.C0();
              _nw144.__ctor(_785_v4, new BigNumber(741), new BigNumber((_dafny.MultiSet.fromElements(_887_cf17)).cardinality()));
              _896_v75 = _nw144;
              let _897_v76;
              _897_v76 = _dafny.Map.Empty.slice().updateUnsafe(_785_v4,_896_v75);
              let _898_v77;
              let _nw145 = Array((new BigNumber(26)).toNumber());
              _nw145[(_dafny.ZERO).toNumber()] = _783_v2;
              _nw145[(_dafny.ONE).toNumber()] = _783_v2;
              _nw145[(new BigNumber(2)).toNumber()] = !(_dafny.Seq.IsPrefixOf((_889_v68).f9, (((_892_v71).contains(_887_cf17)) ? ((_892_v71).get(_887_cf17)) : (_dafny.Seq.UnicodeFromString("ejyl")))));
              _nw145[(new BigNumber(3)).toNumber()] = _783_v2;
              _nw145[(new BigNumber(4)).toNumber()] = false;
              _nw145[(new BigNumber(5)).toNumber()] = (new BigNumber(-50)).isLessThanOrEqualTo((_this).f8);
              _nw145[(new BigNumber(6)).toNumber()] = (_module.D5.create_DC12(_783_v2)).dtor_cf19;
              _nw145[(new BigNumber(7)).toNumber()] = (_887_cf17) && (_887_cf17);
              _nw145[(new BigNumber(8)).toNumber()] = (_783_v2) === (_783_v2);
              _nw145[(new BigNumber(9)).toNumber()] = (_893_v72).contains(false);
              _nw145[(new BigNumber(10)).toNumber()] = false;
              _nw145[(new BigNumber(11)).toNumber()] = !(!_dafny.Seq.contains(_782_v1, new _dafny.CodePoint('q'.codePointAt(0))));
              _nw145[(new BigNumber(12)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_895_v74, _895_v74);
              _nw145[(new BigNumber(13)).toNumber()] = _783_v2;
              _nw145[(new BigNumber(14)).toNumber()] = (((_865_v56).contains((_789_v8).f7)) ? ((_865_v56).get((_789_v8).f7)) : ((((_865_v56).contains(new BigNumber(-911))) ? ((_865_v56).get(new BigNumber(-911))) : (_783_v2))));
              _nw145[(new BigNumber(15)).toNumber()] = _887_cf17;
              _nw145[(new BigNumber(16)).toNumber()] = _783_v2;
              _nw145[(new BigNumber(17)).toNumber()] = _783_v2;
              _nw145[(new BigNumber(18)).toNumber()] = (_module.__default.fm32(globalState)).dtor_cf20;
              _nw145[(new BigNumber(19)).toNumber()] = _887_cf17;
              _nw145[(new BigNumber(20)).toNumber()] = !(_887_cf17);
              _nw145[(new BigNumber(21)).toNumber()] = ((_789_v8).f7).isLessThanOrEqualTo(new BigNumber(391));
              _nw145[(new BigNumber(22)).toNumber()] = _887_cf17;
              _nw145[(new BigNumber(23)).toNumber()] = (_897_v76).contains(_785_v4);
              _nw145[(new BigNumber(24)).toNumber()] = _783_v2;
              _nw145[(new BigNumber(25)).toNumber()] = _887_cf17;
              _898_v77 = _nw145;
              _898_v77 = _898_v77;
              (globalState).f4 = _783_v2;
              let _899_v78;
              _899_v78 = _dafny.Set.fromElements(_782_v1);
              let _index156 = _module.__default.safeIndex(new BigNumber(432), new BigNumber((_898_v77).length));
              (_898_v77)[_index156] = (_899_v78).contains(_dafny.Seq.UnicodeFromString("rhndgwb"));
              let _index157 = _module.__default.safeIndex(new BigNumber(814), new BigNumber((_785_v4).length));
              (_785_v4)[_index157] = _module.__default.safeModuloInt((_dafny.ZERO).minus((((_786_v5).contains((_789_v8).f8)) ? ((_786_v5).get((_789_v8).f8)) : (new BigNumber(718)))), _888_cf16);
              let _index158 = _module.__default.safeIndex(new BigNumber(432), new BigNumber((_898_v77).length));
              let _index159 = _module.__default.safeIndex(new BigNumber(814), new BigNumber((_785_v4).length));
              let _rhs150 = !((_module.__default.safeDivisionInt((_789_v8).f7, (_789_v8).f8)).isLessThan((_789_v8).f8));
              let _rhs151 = _module.__default.fm5(globalState);
              let _rhs152 = (_789_v8).f8;
              let _rhs153 = (_789_v8).f8;
              let _rhs154 = (_888_cf16).isEqualTo((_789_v8).f7);
              let _lhs115 = _898_v77;
              let _lhs116 = _module.__default.safeIndex(new BigNumber(432), new BigNumber((_898_v77).length));
              let _lhs117 = _785_v4;
              let _lhs118 = _module.__default.safeIndex(new BigNumber(814), new BigNumber((_785_v4).length));
              let _lhs119 = globalState;
              _lhs115[_lhs116] = _rhs150;
              _828_v34 = _rhs151;
              _888_cf16 = _rhs152;
              _lhs117[_lhs118] = _rhs153;
              _lhs119.f4 = _rhs154;
            } else {
              let _900___mcc_h11 = (_source19).cf11;
              let _901_cf11 = _900___mcc_h11;
              let _902_v79;
              _902_v79 = _dafny.Set.fromElements((_789_v8).f8);
              let _903_v80;
              _903_v80 = _dafny.Map.Empty.slice().updateUnsafe((_782_v1)[_module.__default.safeIndex((_this).f7, new BigNumber((_782_v1).length))],_902_v79);
              _902_v79 = (((_903_v80).contains(_828_v34)) ? ((_903_v80).get(_828_v34)) : (_902_v79));
              _781_v0 = ((_789_v8).f7).minus((_789_v8).f7);
              let _904_v81;
              _904_v81 = _dafny.MultiSet.fromElements(_828_v34);
              (globalState).f4 = ((_783_v2) ? (_783_v2) : (!(_904_v81).contains(_828_v34)));
              let _905_v82;
              let _nw146 = Array((new BigNumber(12)).toNumber()).fill(false);
              _905_v82 = _nw146;
              let _906_v83;
              let _nw147 = Array((new BigNumber(4)).toNumber());
              _nw147[(_dafny.ZERO).toNumber()] = _905_v82;
              _nw147[(_dafny.ONE).toNumber()] = ((_783_v2) ? (_905_v82) : (_905_v82));
              _nw147[(new BigNumber(2)).toNumber()] = _905_v82;
              _nw147[(new BigNumber(3)).toNumber()] = _905_v82;
              _906_v83 = _nw147;
              let _index160 = _module.__default.safeIndex(new BigNumber(381), new BigNumber((_906_v83).length));
              (_906_v83)[_index160] = _905_v82;
              let _907_v84;
              _907_v84 = _dafny.Map.Empty.slice().updateUnsafe(!(_783_v2),_783_v2);
              let _index161 = _module.__default.safeIndex(new BigNumber(381), new BigNumber((_906_v83).length));
              let _rhs155 = _785_v4;
              let _rhs156 = new BigNumber(151);
              let _rhs157 = _905_v82;
              let _rhs158 = _781_v0;
              let _rhs159 = (((_907_v84).contains(_783_v2)) ? ((_907_v84).get(_783_v2)) : (_783_v2));
              let _lhs120 = _906_v83;
              let _lhs121 = _module.__default.safeIndex(new BigNumber(381), new BigNumber((_906_v83).length));
              let _lhs122 = globalState;
              _785_v4 = _rhs155;
              _781_v0 = _rhs156;
              _lhs120[_lhs121] = _rhs157;
              _781_v0 = _rhs158;
              _lhs122.f4 = _rhs159;
            }
          }
        }
      }
      let _908_v85;
      _908_v85 = new _dafny.CodePoint('s'.codePointAt(0));
      let _909_v86;
      _909_v86 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(276)), ((_910_v2) => function (_911_i10) {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_910_v2,_910_v2)).length);
      })(_783_v2)), _788_v7),_908_v85);
      _909_v86 = (_909_v86).update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-981)), ((_912_v8) => function (_913_i11) {
        return (_912_v8).f7;
      })(_789_v8)), _908_v85);
      let _914_v87;
      _914_v87 = _dafny.Map.Empty.slice().updateUnsafe(_781_v0,(_789_v8).f8);
      _914_v87 = (_914_v87).update((_this).f8, (new BigNumber(-460)).multipliedBy((_789_v8).f7));
      let _915_v88;
      _915_v88 = _dafny.Map.Empty.slice().updateUnsafe(_783_v2,_914_v87);
      let _916_v89;
      _916_v89 = _dafny.Seq.of(false, _783_v2, _783_v2);
      _915_v88 = (_915_v88).update((_916_v89)[_module.__default.safeIndex((_this).f7, new BigNumber((_916_v89).length))], (_914_v87).Merge(_914_v87));
      return;
    }
    m2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = _dafny.ZERO;
      let r2 = false;
      if (p3) {
        (globalState).f4 = true;
        let _917_v0;
        _917_v0 = _dafny.Seq.UnicodeFromString("bjrtfrpd");
        _917_v0 = _dafny.Seq.Concat(_917_v0, _917_v0);
        (globalState).f4 = !((_this).f7).isEqualTo(p0);
        let _918_v1;
        _918_v1 = new _dafny.CodePoint('q'.codePointAt(0));
        let _919_v2;
        _919_v2 = _module.D5.create_DC14(_module.D5.create_DC13(p3, _918_v1));
        _919_v2 = ((p3) ? (_919_v2) : (_919_v2));
        let _920_v3;
        let _init22 = function (_921_i0) {
          return _module.__default.safeModuloInt(_921_i0, (_this).f7);
        };
        let _nw148 = Array((new BigNumber(16)).toNumber());
        for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw148.length); _i0_22++) {
          _nw148[_i0_22] = _init22(new BigNumber(_i0_22));
        }
        _920_v3 = _nw148;
        let _index162 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_920_v3).length));
        (_920_v3)[_index162] = p2;
        let _index163 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_920_v3).length));
        (_920_v3)[_index163] = (_this).f7;
      } else {
        let _922_v4;
        _922_v4 = _dafny.MultiSet.fromElements(p3);
        let _923_v5;
        _923_v5 = _dafny.Map.Empty.slice().updateUnsafe(_922_v4,p2);
        let _924_v6;
        _924_v6 = _dafny.MultiSet.fromElements(new BigNumber(-253));
        let _925_v7;
        _925_v7 = _dafny.Seq.of(new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)));
        let _926_v8;
        _926_v8 = _dafny.Seq.of(_module.__default.fm24(p3, (_dafny.ZERO).minus(new BigNumber((_922_v4).cardinality())), _925_v7, globalState));
        let _927_v9;
        _927_v9 = _module.D6.create_DC16((_923_v5).update(_922_v4, new BigNumber((_924_v6).cardinality())), _926_v8);
        let _source20 = _927_v9;
        if (_source20.is_DC16) {
          let _928___mcc_h0 = (_source20).cf24;
          let _929___mcc_h1 = (_source20).cf25;
          let _930_cf25 = _929___mcc_h1;
          let _931_cf24 = _928___mcc_h0;
          let _932_v10;
          _932_v10 = new _dafny.CodePoint('o'.codePointAt(0));
          let _933_v11;
          _933_v11 = _module.D7.create_DC20(_932_v10, (_this).f7);
          r1 = (_933_v11).dtor_cf33;
          let _934_v12;
          let _nw149 = new _module.C4();
          _nw149.__ctor(p1, p0, (_this).f7);
          _934_v12 = _nw149;
          let _935_v13;
          _935_v13 = _dafny.MultiSet.fromElements(_934_v12);
          let _936_v14;
          _936_v14 = _dafny.Map.Empty.slice().updateUnsafe(false,p1);
          let _937_v15;
          _937_v15 = _dafny.Map.Empty.slice().updateUnsafe((_934_v12).f8,p0);
          let _938_v16;
          _938_v16 = _dafny.Map.Empty.slice().updateUnsafe(_925_v7,(((_922_v4).contains(p1)) ? ((_922_v4).get(p1)) : (p2)));
          let _939_v17;
          _939_v17 = _dafny.Seq.of(new BigNumber((_930_cf25).length), (((_938_v16).contains(_925_v7)) ? ((_938_v16).get(_925_v7)) : ((_934_v12).f7)), (_934_v12).f7, (_934_v12).f7);
          let _940_v18;
          _940_v18 = _dafny.Set.fromElements(true, p3, p3);
          let _941_v19;
          let _nw150 = Array((new BigNumber(18)).toNumber());
          _nw150[(_dafny.ZERO).toNumber()] = new BigNumber(886);
          _nw150[(_dafny.ONE).toNumber()] = (_this).f8;
          _nw150[(new BigNumber(2)).toNumber()] = _module.__default.safeDivisionInt(p2, (_dafny.ZERO).minus(p2));
          _nw150[(new BigNumber(3)).toNumber()] = p0;
          _nw150[(new BigNumber(4)).toNumber()] = (_this).f7;
          _nw150[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("yiww")).length);
          _nw150[(new BigNumber(6)).toNumber()] = new BigNumber((_935_v13).cardinality());
          _nw150[(new BigNumber(7)).toNumber()] = (_this).f7;
          _nw150[(new BigNumber(8)).toNumber()] = ((_934_v12).f7).minus(new BigNumber(636));
          _nw150[(new BigNumber(9)).toNumber()] = (_this).f8;
          _nw150[(new BigNumber(10)).toNumber()] = (_934_v12).f7;
          _nw150[(new BigNumber(11)).toNumber()] = (_this).f7;
          _nw150[(new BigNumber(12)).toNumber()] = (_934_v12).f8;
          _nw150[(new BigNumber(13)).toNumber()] = p0;
          _nw150[(new BigNumber(14)).toNumber()] = (new BigNumber((_936_v14).length)).plus(p0);
          _nw150[(new BigNumber(15)).toNumber()] = ((((_922_v4).contains(!(p1))) ? ((_922_v4).get(!(p1))) : ((((_937_v15).contains((_934_v12).f7)) ? ((_937_v15).get((_934_v12).f7)) : ((_this).f8))))).plus((_939_v17)[_module.__default.safeIndex(new BigNumber((_940_v18).length), new BigNumber((_939_v17).length))]);
          _nw150[(new BigNumber(16)).toNumber()] = ((p1) ? ((_this).f8) : (new BigNumber(-589)));
          _nw150[(new BigNumber(17)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_930_cf25).length), new BigNumber(755));
          _941_v19 = _nw150;
          let _index164 = _module.__default.safeIndex(new BigNumber(248), new BigNumber((_941_v19).length));
          (_941_v19)[_index164] = _module.__default.safeDivisionInt(p0, p2);
          let _index165 = _module.__default.safeIndex(new BigNumber(248), new BigNumber((_941_v19).length));
          (_941_v19)[_index165] = (_this).f8;
          let _942_v20;
          let _943_v21;
          let _944_v22;
          let _out13;
          let _out14;
          let _out15;
          let _outcollector5 = (_this).m3(p3, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(689),new BigNumber((_925_v7).length))).length), _941_v19, (_this).f8, globalState);
          _out13 = _outcollector5[0];
          _out14 = _outcollector5[1];
          _out15 = _outcollector5[2];
          _942_v20 = _out13;
          _943_v21 = _out14;
          _944_v22 = _out15;
          let _945_v23;
          let _946_v24;
          let _947_v25;
          let _out16;
          let _out17;
          let _out18;
          let _outcollector6 = (_this).m3(false, ((_this).f8).minus((_934_v12).f7), _941_v19, (_dafny.ZERO).minus((_this).f7), globalState);
          _out16 = _outcollector6[0];
          _out17 = _outcollector6[1];
          _out18 = _outcollector6[2];
          _945_v23 = _out16;
          _946_v24 = _out17;
          _947_v25 = _out18;
        } else {
          let _948___mcc_h2 = (_source20).cf23;
          let _949_cf23 = _948___mcc_h2;
          r1 = (_this).f7;
          let _950_v26;
          _950_v26 = new _dafny.CodePoint('e'.codePointAt(0));
          let _951_v27;
          _951_v27 = _module.D7.create_DC20(((p3) ? (_950_v26) : (_950_v26)), (_this).f8);
          let _952_v28;
          _952_v28 = _module.D5.create_DC13(p1, _950_v26);
          let _953_v29;
          let _init23 = ((_954_p0) => function (_955_i1) {
            return (_955_i1).multipliedBy(_954_p0);
          })(p0);
          let _nw151 = Array((new BigNumber(8)).toNumber());
          for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw151.length); _i0_23++) {
            _nw151[_i0_23] = _init23(new BigNumber(_i0_23));
          }
          _953_v29 = _nw151;
          let _956_v30;
          _956_v30 = _dafny.Map.Empty.slice().updateUnsafe(_952_v28,_953_v29);
          let _957_v31;
          _957_v31 = _dafny.Seq.of(_956_v30);
          let _rhs160 = ((_957_v31)[_module.__default.safeIndex((_this).f7, new BigNumber((_957_v31).length))]).equals(((_956_v30).update(_952_v28, _953_v29)).update(_952_v28, _953_v29));
          let _rhs161 = ((p3) ? (p1) : (p1));
          let _rhs162 = _module.D7.create_DC20(_950_v26, (_dafny.ZERO).minus(p0));
          let _lhs123 = globalState;
          let _lhs124 = globalState;
          _lhs123.f3 = _rhs160;
          _lhs124.f3 = _rhs161;
          _951_v27 = _rhs162;
          let _index166 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_949_cf23).length));
          (_949_cf23)[_index166] = p3;
          let _index167 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_949_cf23).length));
          (_949_cf23)[_index167] = !(p3);
          r1 = (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(873)), ((_958_v26) => function (_959_i2) {
            return _958_v26;
          })(_950_v26)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(19)), function (_960_i3) {
            return new _dafny.CodePoint('i'.codePointAt(0));
          }))).length)).multipliedBy(new BigNumber(-398));
        }
        let _961_v32;
        _961_v32 = _dafny.Map.Empty.slice().updateUnsafe(p3,new BigNumber(662));
        let _962_v33;
        _962_v33 = _module.D5.create_DC11(_961_v32);
        let _source21 = _962_v33;
        if (_source21.is_DC12) {
          let _963___mcc_h3 = (_source21).cf19;
          let _964_cf19 = _963___mcc_h3;
          let _965_v34;
          _965_v34 = _dafny.Seq.of(_927_v9);
          let _966_v35;
          _966_v35 = _module.D3.create_DC6(_dafny.Set.fromElements(new BigNumber(-311), p0, (_this).f7, (_this).f8, (_this).f8));
          let _967_v36;
          _967_v36 = _dafny.Set.fromElements(p2, (_this).f8);
          let _pat_let_tv14 = _967_v36;
          let _968_v37;
          _968_v37 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(new BigNumber((_965_v34).length), (_this).f8),function (_pat_let12_0) {
            return function (_969_dt__update__tmp_h0) {
              return function (_pat_let13_0) {
                return function (_970_dt__update_hcf9_h0) {
                  return _module.D3.create_DC6(_970_dt__update_hcf9_h0);
                }(_pat_let13_0);
              }(_pat_let_tv14);
            }(_pat_let12_0);
          }(_966_v35));
          _968_v37 = (_968_v37).update(((_dafny.ZERO).minus(p0)).multipliedBy(p0), _966_v35);
          _925_v7 = _925_v7;
          (globalState).f4 = p1;
          r1 = new BigNumber(-788);
        } else if (_source21.is_DC13) {
          let _971___mcc_h4 = (_source21).cf20;
          let _972___mcc_h5 = (_source21).cf21;
          let _973_cf21 = _972___mcc_h5;
          let _974_cf20 = _971___mcc_h4;
          let _975_v38;
          let _nw152 = Array((new BigNumber(8)).toNumber());
          _nw152[(_dafny.ZERO).toNumber()] = ((_this).f7).isLessThan((_this).f7);
          _nw152[(_dafny.ONE).toNumber()] = ((p3) ? (p3) : (_974_cf20));
          _nw152[(new BigNumber(2)).toNumber()] = p1;
          _nw152[(new BigNumber(3)).toNumber()] = true;
          _nw152[(new BigNumber(4)).toNumber()] = _974_cf20;
          _nw152[(new BigNumber(5)).toNumber()] = _974_cf20;
          _nw152[(new BigNumber(6)).toNumber()] = _974_cf20;
          _nw152[(new BigNumber(7)).toNumber()] = _974_cf20;
          _975_v38 = _nw152;
          let _rhs163 = p0;
          let _rhs164 = _975_v38;
          let _rhs165 = (p2).minus(_module.__default.safeDivisionInt(p2, (_this).fm3(_974_cf20, p1, p3, true, globalState)));
          r1 = _rhs163;
          _975_v38 = _rhs164;
          r1 = _rhs165;
          let _976_v39;
          _976_v39 = _dafny.Seq.of(_975_v38);
          let _977_v40;
          let _nw153 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.of());
          _977_v40 = _nw153;
          let _978_v41;
          _978_v41 = _dafny.Seq.of((_this).f8);
          let _index168 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_977_v40).length));
          (_977_v40)[_index168] = _978_v41;
          let _index169 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_977_v40).length));
          let _rhs166 = ((p1) ? (_976_v39) : (_976_v39));
          let _rhs167 = _dafny.Seq.Concat(_dafny.Seq.Concat(_978_v41, _dafny.Seq.of((_this).f7)), _dafny.Seq.Concat(_978_v41, _978_v41));
          let _rhs168 = _925_v7;
          let _rhs169 = _973_cf21;
          let _lhs125 = _977_v40;
          let _lhs126 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_977_v40).length));
          _976_v39 = _rhs166;
          _lhs125[_lhs126] = _rhs167;
          _925_v7 = _rhs168;
          _973_cf21 = _rhs169;
          let _979_v42;
          let _nw154 = new _module.C1();
          _nw154.__ctor((_this).f7, (_dafny.ZERO).minus(p0));
          _979_v42 = _nw154;
          let _980_v43;
          _980_v43 = _dafny.Map.Empty.slice().updateUnsafe(p2,!(p3));
          let _981_v44;
          _981_v44 = _dafny.Set.fromElements((_this).f7, p2, p0, (_dafny.ZERO).minus(new BigNumber((_980_v43).length)));
          let _982_v45;
          _982_v45 = _dafny.Map.Empty.slice().updateUnsafe(_979_v42,new BigNumber((_981_v44).length));
          _982_v45 = (_982_v45).update(_979_v42, p0);
          let _983_v46;
          let _init24 = function (_984_i4) {
            return _module.__default.safeModuloInt(_984_i4, (_this).f8);
          };
          let _nw155 = Array((new BigNumber(23)).toNumber());
          for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw155.length); _i0_24++) {
            _nw155[_i0_24] = _init24(new BigNumber(_i0_24));
          }
          _983_v46 = _nw155;
          let _index170 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_983_v46).length));
          (_983_v46)[_index170] = _module.__default.safeDivisionInt(p2, (_this).f7);
          let _985_v47;
          _985_v47 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,new BigNumber((_925_v7).length));
          let _986_v48;
          _986_v48 = _dafny.Map.Empty.slice().updateUnsafe(p1,_985_v47);
          let _index171 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_983_v46).length));
          (_983_v46)[_index171] = (_module.__default.safeModuloInt(p0, new BigNumber(((((_986_v48).contains(p3)) ? ((_986_v48).get(p3)) : (_985_v47))).length))).plus((_this).f8);
        } else if (_source21.is_DC11) {
          let _987___mcc_h6 = (_source21).cf18;
          let _988_cf18 = _987___mcc_h6;
          (globalState).f3 = p3;
          let _989_v49;
          _989_v49 = new _dafny.CodePoint('j'.codePointAt(0));
          let _990_v50;
          _990_v50 = _dafny.MultiSet.fromElements(_989_v49);
          let _991_v51;
          _991_v51 = _dafny.Seq.of(new BigNumber((_990_v50).cardinality()), p0, p2);
          let _992_v52;
          let _nw156 = new _module.C5();
          _nw156.__ctor(_925_v7, (_this).f7, (_991_v51)[_module.__default.safeIndex(p0, new BigNumber((_991_v51).length))], (_this).f8);
          _992_v52 = _nw156;
          r1 = (_this).f7;
          r1 = p2;
        } else {
          let _993___mcc_h7 = (_source21).cf22;
          let _994_cf22 = _993___mcc_h7;
          _961_v32 = (_961_v32).update(true, (_dafny.ZERO).minus((_this).f7));
          r1 = ((_this).f8).plus(((p3) ? ((_this).f8) : ((_this).f8)));
          let _995_v53;
          let _init25 = function (_996_i5) {
            return _module.__default.safeModuloInt(_996_i5, new BigNumber(648));
          };
          let _nw157 = Array((new BigNumber(7)).toNumber());
          for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw157.length); _i0_25++) {
            _nw157[_i0_25] = _init25(new BigNumber(_i0_25));
          }
          _995_v53 = _nw157;
          let _index172 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((_995_v53).length));
          (_995_v53)[_index172] = new BigNumber((_925_v7).length);
          let _index173 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((_995_v53).length));
          (_995_v53)[_index173] = p0;
          let _index174 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((_995_v53).length));
          (_995_v53)[_index174] = new BigNumber(758);
        }
        let _997_v54;
        _997_v54 = _module.D1.create_DC3(_module.__default.fm23((_this).f8, globalState));
        let _source22 = _997_v54;
        if (_source22.is_DC2) {
          let _998_v55;
          let _init26 = ((_999_p1) => function (_1000_i6) {
            return _999_p1;
          })(p1);
          let _nw158 = Array((new BigNumber(15)).toNumber());
          for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw158.length); _i0_26++) {
            _nw158[_i0_26] = _init26(new BigNumber(_i0_26));
          }
          _998_v55 = _nw158;
          let _index175 = _module.__default.safeIndex(new BigNumber(715), new BigNumber((_998_v55).length));
          (_998_v55)[_index175] = p1;
          let _1001_v56;
          let _nw159 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
          _1001_v56 = _nw159;
          let _index176 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_1001_v56).length));
          (_1001_v56)[_index176] = p2;
          let _1002_v57;
          let _nw160 = new _module.C1();
          _nw160.__ctor((_this).f7, (_this).f8);
          _1002_v57 = _nw160;
          let _1003_v58;
          _1003_v58 = _dafny.Seq.of(_1002_v57);
          let _index177 = _module.__default.safeIndex(new BigNumber(715), new BigNumber((_998_v55).length));
          let _index178 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_1001_v56).length));
          let _rhs170 = (p1) && (!_dafny.areEqual(_1003_v58, _1003_v58));
          let _rhs171 = (p0).isLessThan(new BigNumber((_926_v8).length));
          let _rhs172 = _module.__default.fm16(!(p1), globalState);
          let _rhs173 = (p2).plus(new BigNumber(738));
          let _lhs127 = _998_v55;
          let _lhs128 = _module.__default.safeIndex(new BigNumber(715), new BigNumber((_998_v55).length));
          let _lhs129 = globalState;
          let _lhs130 = _1001_v56;
          let _lhs131 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_1001_v56).length));
          _lhs127[_lhs128] = _rhs170;
          _lhs129.f4 = _rhs171;
          r1 = _rhs172;
          _lhs130[_lhs131] = _rhs173;
          let _1004_v59;
          _1004_v59 = _dafny.Map.Empty.slice().updateUnsafe(((((_922_v4).contains((_998_v55)[_module.__default.safeIndex(new BigNumber(715), new BigNumber((_998_v55).length))])) ? ((_922_v4).get((_998_v55)[_module.__default.safeIndex(new BigNumber(715), new BigNumber((_998_v55).length))])) : ((_this).f8))).minus((_this).f7),p0);
          _1004_v59 = (_1004_v59).update(_module.__default.safeDivisionInt(new BigNumber(594), new BigNumber(147)), p2);
          let _index179 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_1001_v56).length));
          (_1001_v56)[_index179] = new BigNumber((_dafny.Seq.Concat(_925_v7, (((_998_v55)[_module.__default.safeIndex(new BigNumber(715), new BigNumber((_998_v55).length))]) ? (_925_v7) : (_925_v7)))).length);
          let _index180 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_1001_v56).length));
          (_1001_v56)[_index180] = (p2).multipliedBy((_1001_v56)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_1001_v56).length))]);
        } else if (_source22.is_DC1) {
          let _1005___mcc_h8 = (_source22).cf1;
          let _1006_cf1 = _1005___mcc_h8;
          let _1007_v60;
          _1007_v60 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,p1);
          let _1008_v61;
          _1008_v61 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p2,p3)).length));
          let _1009_v62;
          _1009_v62 = _dafny.Seq.of(_1008_v61, _1008_v61);
          let _1010_v63;
          _1010_v63 = _dafny.Map.Empty.slice().updateUnsafe(_925_v7,p3);
          let _1011_v64;
          let _nw161 = Array((new BigNumber(27)).toNumber());
          _nw161[(_dafny.ZERO).toNumber()] = (_this).f8;
          _nw161[(_dafny.ONE).toNumber()] = new BigNumber((_1007_v60).length);
          _nw161[(new BigNumber(2)).toNumber()] = _module.__default.fm16(p1, globalState);
          _nw161[(new BigNumber(3)).toNumber()] = _module.__default.fm20(p3, p1, p1, globalState);
          _nw161[(new BigNumber(4)).toNumber()] = (_this).f7;
          _nw161[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus((_this).f7);
          _nw161[(new BigNumber(6)).toNumber()] = (_this).f7;
          _nw161[(new BigNumber(7)).toNumber()] = p0;
          _nw161[(new BigNumber(8)).toNumber()] = new BigNumber((_925_v7).length);
          _nw161[(new BigNumber(9)).toNumber()] = new BigNumber((_1009_v62).length);
          _nw161[(new BigNumber(10)).toNumber()] = new BigNumber((_926_v8).length);
          _nw161[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus((_this).f8);
          _nw161[(new BigNumber(12)).toNumber()] = new BigNumber(500);
          _nw161[(new BigNumber(13)).toNumber()] = p0;
          _nw161[(new BigNumber(14)).toNumber()] = p2;
          _nw161[(new BigNumber(15)).toNumber()] = new BigNumber(532);
          _nw161[(new BigNumber(16)).toNumber()] = p2;
          _nw161[(new BigNumber(17)).toNumber()] = p2;
          _nw161[(new BigNumber(18)).toNumber()] = new BigNumber((_925_v7).length);
          _nw161[(new BigNumber(19)).toNumber()] = p0;
          _nw161[(new BigNumber(20)).toNumber()] = (_this).f7;
          _nw161[(new BigNumber(21)).toNumber()] = (_this).f8;
          _nw161[(new BigNumber(22)).toNumber()] = (_this).f8;
          _nw161[(new BigNumber(23)).toNumber()] = new BigNumber((_1010_v63).length);
          _nw161[(new BigNumber(24)).toNumber()] = new BigNumber(274);
          _nw161[(new BigNumber(25)).toNumber()] = p2;
          _nw161[(new BigNumber(26)).toNumber()] = p0;
          _1011_v64 = _nw161;
          let _1012_v65;
          _1012_v65 = _dafny.Seq.of(_1011_v64, _1011_v64, _1011_v64, _1011_v64, _1011_v64);
          let _1013_v66;
          let _nw162 = Array((new BigNumber(10)).toNumber());
          _nw162[(_dafny.ZERO).toNumber()] = _1011_v64;
          _nw162[(_dafny.ONE).toNumber()] = _1011_v64;
          _nw162[(new BigNumber(2)).toNumber()] = _1011_v64;
          _nw162[(new BigNumber(3)).toNumber()] = _1011_v64;
          _nw162[(new BigNumber(4)).toNumber()] = _1011_v64;
          _nw162[(new BigNumber(5)).toNumber()] = _1011_v64;
          _nw162[(new BigNumber(6)).toNumber()] = (_1012_v65)[_module.__default.safeIndex((_dafny.ZERO).minus(p2), new BigNumber((_1012_v65).length))];
          _nw162[(new BigNumber(7)).toNumber()] = _1011_v64;
          _nw162[(new BigNumber(8)).toNumber()] = _1011_v64;
          _nw162[(new BigNumber(9)).toNumber()] = _1011_v64;
          _1013_v66 = _nw162;
          let _index181 = _module.__default.safeIndex(new BigNumber(861), new BigNumber((_1013_v66).length));
          (_1013_v66)[_index181] = _1011_v64;
          let _index182 = _module.__default.safeIndex(new BigNumber(861), new BigNumber((_1013_v66).length));
          (_1013_v66)[_index182] = _1011_v64;
          let _1014_v67;
          _1014_v67 = _dafny.Set.fromElements(_925_v7);
          (globalState).f4 = ((_1014_v67).Difference(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("hwlofrqbi")))).IsSubsetOf((((_926_v8)[_module.__default.safeIndex(p0, new BigNumber((_926_v8).length))]) ? (_1014_v67) : (_1014_v67)));
          let _1015_v68;
          _1015_v68 = _dafny.Map.Empty.slice().updateUnsafe(p1,_925_v7);
          let _index183 = _module.__default.safeIndex(new BigNumber(155), new BigNumber((_1011_v64).length));
          (_1011_v64)[_index183] = new BigNumber((_1015_v68).length);
          let _1016_v69;
          _1016_v69 = _dafny.Seq.of(p0);
          let _1017_v70;
          _1017_v70 = _dafny.Map.Empty.slice().updateUnsafe(_1006_cf1,_925_v7);
          let _1018_v71;
          _1018_v71 = _dafny.Map.Empty.slice().updateUnsafe(p3,((p3) ? (_1013_v66) : (_1013_v66)));
          let _index184 = _module.__default.safeIndex(new BigNumber(155), new BigNumber((_1011_v64).length));
          let _rhs174 = ((_dafny.ZERO).minus((_this).f8)).plus((_this).f7);
          let _rhs175 = _1016_v69;
          let _rhs176 = ((!_dafny.Seq.contains(_925_v7, _1006_cf1)) ? (p1) : ((((_1007_v60).contains((_this).f8)) ? ((_1007_v60).get((_this).f8)) : (!(false)))));
          let _rhs177 = _module.__default.safeDivisionInt((_this).f7, new BigNumber((_1017_v70).length));
          let _rhs178 = (((_1018_v71).contains(true)) ? ((_1018_v71).get(true)) : (_1013_v66));
          let _lhs132 = _1011_v64;
          let _lhs133 = _module.__default.safeIndex(new BigNumber(155), new BigNumber((_1011_v64).length));
          let _lhs134 = globalState;
          _lhs132[_lhs133] = _rhs174;
          _1016_v69 = _rhs175;
          _lhs134.f3 = _rhs176;
          r1 = _rhs177;
          _1013_v66 = _rhs178;
          let _index185 = _module.__default.safeIndex(new BigNumber(155), new BigNumber((_1011_v64).length));
          (_1011_v64)[_index185] = ((_dafny.ZERO).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus((_this).f7), (_this).f8))).multipliedBy(new BigNumber(646));
        } else {
          let _1019___mcc_h9 = (_source22).cf2;
          let _1020_cf2 = _1019___mcc_h9;
          let _1021_v72;
          _1021_v72 = _dafny.Set.fromElements(p2);
          _1021_v72 = _1021_v72;
          let _1022_v73;
          let _nw163 = new _module.C5();
          _nw163.__ctor(_dafny.Seq.UnicodeFromString("svx"), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(532)), function (_1023_i7) {
            return new _dafny.CodePoint('f'.codePointAt(0));
          })).length), new BigNumber(135), (_this).f8);
          _1022_v73 = _nw163;
          let _1024_v74;
          _1024_v74 = _dafny.Map.Empty.slice().updateUnsafe(_925_v7,_1022_v73);
          let _1025_v75;
          _1025_v75 = _dafny.Seq.of((((_1024_v74).contains((_this).fm2(p3, globalState))) ? ((_1024_v74).get((_this).fm2(p3, globalState))) : (_1022_v73)), _1022_v73);
          let _1026_v76;
          let _nw164 = new _module.C2();
          _nw164.__ctor(new BigNumber((_1025_v75).length), _module.__default.fm20(p3, false, p1, globalState));
          _1026_v76 = _nw164;
          _926_v8 = _926_v8;
          let _1027_v77;
          let _nw165 = Array((new BigNumber(3)).toNumber()).fill([]);
          _1027_v77 = _nw165;
          let _1028_v78;
          let _nw166 = Array((new BigNumber(6)).toNumber()).fill([]);
          _1028_v78 = _nw166;
          let _index186 = _module.__default.safeIndex(new BigNumber(418), new BigNumber((_1027_v77).length));
          (_1027_v77)[_index186] = _1028_v78;
          let _1029_v79;
          let _nw167 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.of());
          _1029_v79 = _nw167;
          let _1030_v80;
          _1030_v80 = _module.D7.create_DC19(p0, p1, (_this).f7);
          let _1031_v81;
          _1031_v81 = _dafny.Map.Empty.slice().updateUnsafe(((_1022_v73).f8).minus(p2),((_1030_v80).dtor_cf31).isLessThanOrEqualTo(p2));
          let _index187 = _module.__default.safeIndex(new BigNumber(418), new BigNumber((_1027_v77).length));
          let _rhs179 = _1028_v78;
          let _rhs180 = _925_v7;
          let _rhs181 = _1029_v79;
          let _rhs182 = p3;
          let _rhs183 = _1031_v81;
          let _lhs135 = _1027_v77;
          let _lhs136 = _module.__default.safeIndex(new BigNumber(418), new BigNumber((_1027_v77).length));
          let _lhs137 = globalState;
          _lhs135[_lhs136] = _rhs179;
          _925_v7 = _rhs180;
          _1029_v79 = _rhs181;
          _lhs137.f3 = _rhs182;
          _1031_v81 = _rhs183;
        }
        r1 = p2;
        let _1032_v82;
        let _nw168 = Array((new BigNumber(15)).toNumber()).fill([]);
        _1032_v82 = _nw168;
        let _1033_v83;
        let _nw169 = Array((new BigNumber(19)).toNumber()).fill(false);
        _1033_v83 = _nw169;
        let _1034_v84;
        _1034_v84 = _dafny.Map.Empty.slice().updateUnsafe(p3,_1033_v83);
        let _1035_v85;
        _1035_v85 = new _dafny.CodePoint('w'.codePointAt(0));
        let _1036_v86;
        let _nw170 = Array((new BigNumber(14)).toNumber());
        _nw170[(_dafny.ZERO).toNumber()] = (((_924_v6).contains(p0)) ? ((_924_v6).get(p0)) : ((_this).f8));
        _nw170[(_dafny.ONE).toNumber()] = (_this).f7;
        _nw170[(new BigNumber(2)).toNumber()] = p0;
        _nw170[(new BigNumber(3)).toNumber()] = new BigNumber((_1034_v84).length);
        _nw170[(new BigNumber(4)).toNumber()] = (_this).f7;
        _nw170[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(p0);
        _nw170[(new BigNumber(6)).toNumber()] = p2;
        _nw170[(new BigNumber(7)).toNumber()] = (_this).fm4(true, true, globalState);
        _nw170[(new BigNumber(8)).toNumber()] = p2;
        _nw170[(new BigNumber(9)).toNumber()] = _module.__default.fm20(p1, p1, p3, globalState);
        _nw170[(new BigNumber(10)).toNumber()] = (_this).f7;
        _nw170[(new BigNumber(11)).toNumber()] = new BigNumber(965);
        _nw170[(new BigNumber(12)).toNumber()] = new BigNumber(((_dafny.MultiSet.fromElements(_1035_v85)).update(_1035_v85, _module.__default.abs(new BigNumber((_925_v7).length)))).cardinality());
        _nw170[(new BigNumber(13)).toNumber()] = p0;
        _1036_v86 = _nw170;
        let _index188 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_1032_v82).length));
        (_1032_v82)[_index188] = _1036_v86;
        let _1037_v87;
        _1037_v87 = _module.D2.create_DC4(_922_v4);
        let _1038_v88;
        _1038_v88 = _dafny.MultiSet.fromElements(_1037_v87, _1037_v87, _1037_v87);
        let _1039_v89;
        _1039_v89 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,p3);
        let _index189 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_1032_v82).length));
        let _rhs184 = ((p1) ? (new BigNumber(-197)) : (_module.__default.safeDivisionInt((_this).f8, new BigNumber((_1039_v89).length))));
        let _rhs185 = _1036_v86;
        let _rhs186 = (_dafny.MultiSet.fromElements(_1037_v87)).Union(_1038_v88);
        let _rhs187 = ((_this).f8).plus((_this).f7);
        let _lhs138 = _1032_v82;
        let _lhs139 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_1032_v82).length));
        r1 = _rhs184;
        _lhs138[_lhs139] = _rhs185;
        _1038_v88 = _rhs186;
        r1 = _rhs187;
      }
      if (p3) {
        let _1040_v90;
        _1040_v90 = new _dafny.CodePoint('h'.codePointAt(0));
        _1040_v90 = _1040_v90;
        r1 = ((!(true)) ? ((_this).f7) : (((p1) ? (p0) : ((_this).f7))));
        let _1041_v91;
        let _init27 = function (_1042_i8) {
          return (_1042_i8).multipliedBy(new BigNumber(-317));
        };
        let _nw171 = Array((new BigNumber(8)).toNumber());
        for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw171.length); _i0_27++) {
          _nw171[_i0_27] = _init27(new BigNumber(_i0_27));
        }
        _1041_v91 = _nw171;
        let _index190 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_1041_v91).length));
        (_1041_v91)[_index190] = new BigNumber(629);
        let _1043_v92;
        _1043_v92 = _dafny.MultiSet.fromElements(p2, p2);
        let _1044_v94;
        _1044_v94 = _dafny.Set.fromElements(_module.__default.fm29(p3, _1043_v92, globalState), (function () {
          let _coll38 = new _dafny.Set();
          for (const _compr_38 of _dafny.IntegerRange(new BigNumber(224), new BigNumber(478))) {
            let _1045_v93 = _compr_38;
            if (((new BigNumber(224)).isLessThanOrEqualTo(_1045_v93)) && ((_1045_v93).isLessThan(new BigNumber(478)))) {
              _coll38.add((_1045_v93).plus(new BigNumber(701)));
            }
          }
          return _coll38;
        }()).Intersect(_dafny.Set.fromElements(new BigNumber(-460), (_this).f8, (_dafny.ZERO).minus(p2))));
        let _1046_v95;
        _1046_v95 = _dafny.Seq.UnicodeFromString("ctt");
        let _1047_v96;
        _1047_v96 = _dafny.Set.fromElements(new BigNumber(788));
        let _index191 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_1041_v91).length));
        let _rhs188 = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_this).f7, new BigNumber((_1046_v95).length)));
        let _rhs189 = ((_1044_v94).Difference(_dafny.Set.fromElements(_1047_v96))).Difference(_1044_v94);
        let _rhs190 = p0;
        let _lhs140 = _1041_v91;
        let _lhs141 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_1041_v91).length));
        _lhs140[_lhs141] = _rhs188;
        _1044_v94 = _rhs189;
        r1 = _rhs190;
        r1 = ((_this).fm3(p3, p1, p3, false, globalState)).multipliedBy((new BigNumber(-226)).multipliedBy((_1041_v91)[_module.__default.safeIndex(new BigNumber(552), new BigNumber((_1041_v91).length))]));
        let _1048_v97;
        _1048_v97 = _dafny.Seq.of(p1, (_module.D5.create_DC13(_module.__default.fm24(p1, (_this).f7, _1046_v95, globalState), _1040_v90)).dtor_cf20, !(p3));
        let _1049_v98;
        let _nw172 = new _module.C2();
        _nw172.__ctor((_this).f8, p2);
        _1049_v98 = _nw172;
        let _1050_v99;
        _1050_v99 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1047_v96).length),_1049_v98);
        let _1051_v100;
        let _nw173 = Array((new BigNumber(4)).toNumber());
        _nw173[(_dafny.ZERO).toNumber()] = p1;
        _nw173[(_dafny.ONE).toNumber()] = (_1048_v97)[_module.__default.safeIndex(new BigNumber((_1050_v99).length), new BigNumber((_1048_v97).length))];
        _nw173[(new BigNumber(2)).toNumber()] = p3;
        _nw173[(new BigNumber(3)).toNumber()] = p3;
        _1051_v100 = _nw173;
        let _index192 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_1051_v100).length));
        (_1051_v100)[_index192] = p1;
        let _index193 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_1051_v100).length));
        (_1051_v100)[_index193] = (_1048_v97)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_1048_v97).length))];
      } else {
        let _1052_v101;
        _1052_v101 = _dafny.Map.Empty.slice().updateUnsafe(false,p1);
        r1 = (new BigNumber(((_1052_v101).update(p3, p1)).length)).multipliedBy((new BigNumber((_dafny.Seq.UnicodeFromString("ymjehvl")).length)).plus((_this).f7));
        let _1053_v102;
        _1053_v102 = _dafny.Seq.UnicodeFromString("uaehms");
        let _1054_v103;
        _1054_v103 = _dafny.Seq.of(_1053_v102, _dafny.Seq.Create(_module.__default.abs(new BigNumber(274)), function (_1055_i9) {
          return new _dafny.CodePoint('u'.codePointAt(0));
        }), _1053_v102);
        if (!_dafny.areEqual(_dafny.Seq.Concat(_1053_v102, _1053_v102), _dafny.Seq.Concat((_1054_v103)[_module.__default.safeIndex((_this).f7, new BigNumber((_1054_v103).length))], _1053_v102))) {
          r1 = (_this).f8;
          let _1056_v104;
          _1056_v104 = _dafny.Seq.of(p1);
          let _1057_v105;
          _1057_v105 = new _dafny.CodePoint('i'.codePointAt(0));
          let _1058_v106;
          _1058_v106 = _dafny.Seq.of((_this).f8);
          let _1059_v107;
          let _nw174 = Array((new BigNumber(10)).toNumber());
          _nw174[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.update(_1056_v104, _module.__default.safeIndex(p2, new BigNumber((_1056_v104).length)), p1)).length);
          _nw174[(_dafny.ONE).toNumber()] = (_module.D7.create_DC20(_1057_v105, new BigNumber((_1054_v103).length))).dtor_cf33;
          _nw174[(new BigNumber(2)).toNumber()] = (_this).f7;
          _nw174[(new BigNumber(3)).toNumber()] = (p2).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(p3),p1)).length));
          _nw174[(new BigNumber(4)).toNumber()] = ((p3) ? (p2) : ((_this).f7));
          _nw174[(new BigNumber(5)).toNumber()] = (new BigNumber((_1058_v106).length)).plus((_dafny.ZERO).minus((_this).f8));
          _nw174[(new BigNumber(6)).toNumber()] = p2;
          _nw174[(new BigNumber(7)).toNumber()] = (new BigNumber(545)).minus((_dafny.ZERO).minus(new BigNumber(-723)));
          _nw174[(new BigNumber(8)).toNumber()] = p2;
          _nw174[(new BigNumber(9)).toNumber()] = (_this).f7;
          _1059_v107 = _nw174;
          let _1060_v108;
          _1060_v108 = _dafny.Map.Empty.slice().updateUnsafe(_1057_v105,p3);
          let _1061_v109;
          let _nw175 = new _module.C5();
          _nw175.__ctor(_1053_v102, (_this).f7, (_this).f7, (_dafny.ZERO).minus(new BigNumber((_1060_v108).length)));
          _1061_v109 = _nw175;
          let _1062_v110;
          _1062_v110 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,_1061_v109);
          let _index194 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_1059_v107).length));
          (_1059_v107)[_index194] = ((_this).f8).minus(new BigNumber(((_1062_v110).update((_this).f8, _1061_v109)).length));
          let _index195 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_1059_v107).length));
          let _rhs191 = (_this).f7;
          let _rhs192 = (_this).f8;
          let _lhs142 = _1059_v107;
          let _lhs143 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_1059_v107).length));
          _lhs142[_lhs143] = _rhs191;
          r1 = _rhs192;
          _1057_v105 = _1057_v105;
          let _1063_v111;
          _1063_v111 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),(_this).f8);
          let _1064_v113;
          _1064_v113 = _dafny.Set.fromElements(_module.__default.fm5(globalState));
          let _1065_v114;
          _1065_v114 = _dafny.Seq.of((_dafny.Set.fromElements(_1057_v105)).Union(function () {
            let _coll39 = new _dafny.Set();
            for (const _compr_39 of (_1063_v111).Keys.Elements) {
              let _1066_v112 = _compr_39;
              if ((_1063_v111).contains(_1066_v112)) {
                _coll39.add(_1066_v112);
              }
            }
            return _coll39;
          }()), _1064_v113);
          _1065_v114 = _dafny.Seq.Concat(_1065_v114, _dafny.Seq.Concat(_1065_v114, _1065_v114));
          let _index196 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_1059_v107).length));
          (_1059_v107)[_index196] = ((_this).f7).plus((_this).f7);
        } else {
          let _1067_v115;
          _1067_v115 = _dafny.Seq.of(p1);
          let _1068_v116;
          _1068_v116 = _module.D2.create_DC5(p1, (_this).f8, new BigNumber((_dafny.Seq.update(_1067_v115, _module.__default.safeIndex(new BigNumber((_1067_v115).length), new BigNumber((_1067_v115).length)), p1)).length), (_this).f7, _module.__default.fm24(p3, new BigNumber(998), _1053_v102, globalState));
          let _1069_v117;
          _1069_v117 = _dafny.Set.fromElements(p3);
          let _1070_v118;
          let _nw176 = Array((new BigNumber(17)).toNumber());
          _nw176[(_dafny.ZERO).toNumber()] = !(p3);
          _nw176[(_dafny.ONE).toNumber()] = p1;
          _nw176[(new BigNumber(2)).toNumber()] = !(!((((_1068_v116).dtor_cf4) ? (p3) : ((_this).fm1(_1069_v117, (_this).f7, new BigNumber(-719), (_dafny.ZERO).minus((_dafny.ZERO).minus(p0)), globalState)))));
          _nw176[(new BigNumber(3)).toNumber()] = p3;
          _nw176[(new BigNumber(4)).toNumber()] = p3;
          _nw176[(new BigNumber(5)).toNumber()] = p3;
          _nw176[(new BigNumber(6)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_1053_v102, _dafny.Seq.Create(_module.__default.abs(new BigNumber(279)), function (_1071_i10) {
            return new _dafny.CodePoint('b'.codePointAt(0));
          }));
          _nw176[(new BigNumber(7)).toNumber()] = p1;
          _nw176[(new BigNumber(8)).toNumber()] = p1;
          _nw176[(new BigNumber(9)).toNumber()] = !(p3) || (true);
          _nw176[(new BigNumber(10)).toNumber()] = p3;
          _nw176[(new BigNumber(11)).toNumber()] = p3;
          _nw176[(new BigNumber(12)).toNumber()] = p1;
          _nw176[(new BigNumber(13)).toNumber()] = !(p1) || (true);
          _nw176[(new BigNumber(14)).toNumber()] = p3;
          _nw176[(new BigNumber(15)).toNumber()] = p3;
          _nw176[(new BigNumber(16)).toNumber()] = p3;
          _1070_v118 = _nw176;
          let _index197 = _module.__default.safeIndex(new BigNumber(271), new BigNumber((_1070_v118).length));
          (_1070_v118)[_index197] = true;
          let _index198 = _module.__default.safeIndex(new BigNumber(271), new BigNumber((_1070_v118).length));
          (_1070_v118)[_index198] = !((_this).f8).isEqualTo(((_this).f8).plus(p2));
          let _1072_v119;
          _1072_v119 = _dafny.Set.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_1052_v101).length))), p2);
          r1 = (p0).multipliedBy(new BigNumber((_1072_v119).length));
          let _1073_v120;
          _1073_v120 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f7);
          let _1074_v121;
          let _nw177 = new _module.C1();
          _nw177.__ctor((((_1073_v120).contains(p0)) ? ((_1073_v120).get(p0)) : ((_this).f7)), ((!(p3)) ? ((_this).f7) : ((_this).fm3(p1, true, (_1070_v118)[_module.__default.safeIndex(new BigNumber(271), new BigNumber((_1070_v118).length))], p1, globalState))));
          _1074_v121 = _nw177;
          let _1075_v122;
          _1075_v122 = _module.D5.create_DC13(p1, new _dafny.CodePoint('q'.codePointAt(0)));
          let _1076_v123;
          _1076_v123 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(_module.__default.fm16(p3, globalState), p0),_dafny.Seq.of(_1075_v122, _1075_v122, _module.__default.fm32(globalState), _module.__default.fm32(globalState)));
          let _1077_v124;
          _1077_v124 = _dafny.MultiSet.fromElements(p3);
          let _1078_v125;
          _1078_v125 = _dafny.MultiSet.fromElements((((_1077_v124).contains(p3)) ? ((_1077_v124).get(p3)) : ((_this).f7)));
          let _1079_v126;
          _1079_v126 = _dafny.Seq.of(_module.__default.safeDivisionInt(new BigNumber((_1078_v125).cardinality()), new BigNumber(827)), (_dafny.ZERO).minus((new BigNumber(-451)).plus(p0)), p0, _module.__default.safeModuloInt((_this).f7, p2));
          let _index199 = _module.__default.safeIndex(new BigNumber(271), new BigNumber((_1070_v118).length));
          let _rhs193 = ((p1) ? (false) : (p1));
          let _rhs194 = _1053_v102;
          let _rhs195 = _1076_v123;
          let _rhs196 = (_1070_v118)[_module.__default.safeIndex(new BigNumber(271), new BigNumber((_1070_v118).length))];
          let _rhs197 = _1079_v126;
          let _lhs144 = globalState;
          let _lhs145 = _1070_v118;
          let _lhs146 = _module.__default.safeIndex(new BigNumber(271), new BigNumber((_1070_v118).length));
          _lhs144.f4 = _rhs193;
          _1053_v102 = _rhs194;
          _1076_v123 = _rhs195;
          _lhs145[_lhs146] = _rhs196;
          _1079_v126 = _rhs197;
          r1 = (new BigNumber((_dafny.Seq.Concat(_1053_v102, _1053_v102)).length)).plus((((_1078_v125).contains(_module.__default.fm20(p3, (_1070_v118)[_module.__default.safeIndex(new BigNumber(271), new BigNumber((_1070_v118).length))], p1, globalState))) ? ((_1078_v125).get(_module.__default.fm20(p3, (_1070_v118)[_module.__default.safeIndex(new BigNumber(271), new BigNumber((_1070_v118).length))], p1, globalState))) : (new BigNumber((_1067_v115).length))));
        }
        let _1080_v127;
        let _nw178 = new _module.C3();
        _nw178.__ctor(!(p3), (_this).f7, p2);
        _1080_v127 = _nw178;
        let _1081_v128;
        let _nw179 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _1081_v128 = _nw179;
        let _index200 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_1081_v128).length));
        (_1081_v128)[_index200] = _1053_v102;
        let _index201 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_1081_v128).length));
        (_1081_v128)[_index201] = _dafny.Seq.Concat(_1053_v102, _1053_v102);
        let _1082_v129;
        _1082_v129 = _dafny.Set.fromElements((_1080_v127).f12);
        let _1083_v130;
        _1083_v130 = _dafny.Seq.of(p1, p3);
        let _1084_v131;
        let _nw180 = Array((new BigNumber(18)).toNumber());
        _nw180[(_dafny.ZERO).toNumber()] = new BigNumber((_1053_v102).length);
        _nw180[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_this).f8, p2));
        _nw180[(new BigNumber(2)).toNumber()] = p0;
        _nw180[(new BigNumber(3)).toNumber()] = p2;
        _nw180[(new BigNumber(4)).toNumber()] = p2;
        _nw180[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus((_this).f8);
        _nw180[(new BigNumber(6)).toNumber()] = (_this).f7;
        _nw180[(new BigNumber(7)).toNumber()] = (_this).f8;
        _nw180[(new BigNumber(8)).toNumber()] = new BigNumber((_1082_v129).length);
        _nw180[(new BigNumber(9)).toNumber()] = new BigNumber((_1083_v130).length);
        _nw180[(new BigNumber(10)).toNumber()] = new BigNumber(-703);
        _nw180[(new BigNumber(11)).toNumber()] = _module.__default.safeModuloInt(p2, new BigNumber((_1082_v129).length));
        _nw180[(new BigNumber(12)).toNumber()] = new BigNumber(((_1080_v127).fm2(true, globalState)).length);
        _nw180[(new BigNumber(13)).toNumber()] = p0;
        _nw180[(new BigNumber(14)).toNumber()] = ((_this).f7).plus(p2);
        _nw180[(new BigNumber(15)).toNumber()] = (_this).f8;
        _nw180[(new BigNumber(16)).toNumber()] = (new BigNumber(125)).minus((_this).f8);
        _nw180[(new BigNumber(17)).toNumber()] = p2;
        _1084_v131 = _nw180;
        _1084_v131 = _1084_v131;
      }
      if ((p2).isLessThanOrEqualTo(p0)) {
        let _1085_v132;
        _1085_v132 = _dafny.Seq.of(p0);
        let _1086_v133;
        _1086_v133 = _dafny.Seq.of(_1085_v132, _dafny.Seq.Create(_module.__default.abs(new BigNumber(774)), function (_1087_i11) {
          return new BigNumber(431);
        }));
        let _1088_v134;
        _1088_v134 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,new BigNumber((_1086_v133).length));
        _1088_v134 = (_1088_v134).update(new BigNumber(-267), ((_this).f8).plus(p0));
        let _1089_v135;
        _1089_v135 = _dafny.MultiSet.fromElements(p1, p3, p3, !(p3), !(p1));
        let _1090_v136;
        _1090_v136 = _dafny.Map.Empty.slice().updateUnsafe(_1089_v135,p1);
        _1090_v136 = (_1090_v136).update(_1089_v135, ((_this).f8).isLessThan((_this).f7));
        let _1091_v137;
        _1091_v137 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,p3);
        _1091_v137 = (_1091_v137).update((_dafny.ZERO).minus((_this).f7), true);
        (globalState).f4 = p3;
        (globalState).f3 = ((_this).f7).isLessThan(p0);
      } else {
        let _1092_v138;
        _1092_v138 = _dafny.Seq.of(p2);
        let _1093_v139;
        _1093_v139 = _dafny.Seq.UnicodeFromString("fxgvl");
        let _1094_v140;
        _1094_v140 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1092_v138).length),new BigNumber((_1093_v139).length));
        let _1095_v141;
        _1095_v141 = _dafny.Set.fromElements(p1, p1, p1, p3, p3);
        let _1096_v142;
        _1096_v142 = _dafny.MultiSet.fromElements(_1095_v141);
        let _1097_v143;
        _1097_v143 = _dafny.Map.Empty.slice().updateUnsafe(!(false),true);
        let _1098_v144;
        _1098_v144 = _dafny.MultiSet.fromElements(p2);
        let _1099_v145;
        _1099_v145 = _dafny.MultiSet.fromElements(_1098_v144, _1098_v144);
        let _1100_v146;
        let _nw181 = Array((new BigNumber(9)).toNumber());
        _nw181[(_dafny.ZERO).toNumber()] = (_this).f7;
        _nw181[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus((_this).f8);
        _nw181[(new BigNumber(2)).toNumber()] = (_this).f7;
        _nw181[(new BigNumber(3)).toNumber()] = new BigNumber(34);
        _nw181[(new BigNumber(4)).toNumber()] = (((_1099_v145).contains(_1098_v144)) ? ((_1099_v145).get(_1098_v144)) : (p0));
        _nw181[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,!(p3))).length);
        _nw181[(new BigNumber(6)).toNumber()] = p0;
        _nw181[(new BigNumber(7)).toNumber()] = new BigNumber((_1093_v139).length);
        _nw181[(new BigNumber(8)).toNumber()] = p2;
        _1100_v146 = _nw181;
        let _1101_v147;
        _1101_v147 = _dafny.Map.Empty.slice().updateUnsafe(_1100_v146,p3);
        let _1102_v148;
        let _nw182 = Array((new BigNumber(26)).toNumber());
        _nw182[(_dafny.ZERO).toNumber()] = p2;
        _nw182[(_dafny.ONE).toNumber()] = (((_1094_v140).contains(p0)) ? ((_1094_v140).get(p0)) : (p0));
        _nw182[(new BigNumber(2)).toNumber()] = (_this).f8;
        _nw182[(new BigNumber(3)).toNumber()] = (_this).f8;
        _nw182[(new BigNumber(4)).toNumber()] = new BigNumber(201);
        _nw182[(new BigNumber(5)).toNumber()] = p2;
        _nw182[(new BigNumber(6)).toNumber()] = (((_1096_v142).contains(_1095_v141)) ? ((_1096_v142).get(_1095_v141)) : (new BigNumber((_1097_v143).length)));
        _nw182[(new BigNumber(7)).toNumber()] = p2;
        _nw182[(new BigNumber(8)).toNumber()] = new BigNumber(-434);
        _nw182[(new BigNumber(9)).toNumber()] = (_1092_v138)[_module.__default.safeIndex(p0, new BigNumber((_1092_v138).length))];
        _nw182[(new BigNumber(10)).toNumber()] = (_this).fm3(p3, p1, p1, (((_1097_v143).contains(p3)) ? ((_1097_v143).get(p3)) : (p1)), globalState);
        _nw182[(new BigNumber(11)).toNumber()] = p2;
        _nw182[(new BigNumber(12)).toNumber()] = (_this).f7;
        _nw182[(new BigNumber(13)).toNumber()] = p2;
        _nw182[(new BigNumber(14)).toNumber()] = new BigNumber((_1092_v138).length);
        _nw182[(new BigNumber(15)).toNumber()] = (_this).f7;
        _nw182[(new BigNumber(16)).toNumber()] = (_this).f7;
        _nw182[(new BigNumber(17)).toNumber()] = new BigNumber((_1092_v138).length);
        _nw182[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1101_v147).length));
        _nw182[(new BigNumber(19)).toNumber()] = p0;
        _nw182[(new BigNumber(20)).toNumber()] = p0;
        _nw182[(new BigNumber(21)).toNumber()] = p2;
        _nw182[(new BigNumber(22)).toNumber()] = p2;
        _nw182[(new BigNumber(23)).toNumber()] = _module.__default.fm20(p3, true, false, globalState);
        _nw182[(new BigNumber(24)).toNumber()] = p2;
        _nw182[(new BigNumber(25)).toNumber()] = p0;
        _1102_v148 = _nw182;
        let _1103_v149;
        let _nw183 = new _module.C0();
        _nw183.__ctor(_1100_v146, (_this).f7, (_this).f8);
        _1103_v149 = _nw183;
        let _1104_v150;
        _1104_v150 = _dafny.Seq.of(p1);
        r0 = _dafny.Seq.Concat(_1104_v150, _1104_v150);
        let _1105_v151;
        _1105_v151 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f7).isLessThan((_this).f7),(_this).f7);
        _1105_v151 = (_1105_v151).update(p3, (p2).minus((_this).f8));
        r1 = p2;
        let _arr14 = _1103_v149.f13;
        let _index202 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_1103_v149.f13).length));
        _arr14[_index202] = p2;
        let _arr15 = _1103_v149.f13;
        let _index203 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_1103_v149.f13).length));
        _arr15[_index203] = p2;
      }
      if (p1) {
        let _1106_v152;
        _1106_v152 = new _dafny.CodePoint('e'.codePointAt(0));
        let _1107_v153;
        _1107_v153 = _dafny.Seq.UnicodeFromString("vlqsggw");
        let _1108_v154;
        _1108_v154 = _dafny.Seq.of((_1107_v153)[_module.__default.safeIndex(p0, new BigNumber((_1107_v153).length))], _1106_v152, _module.__default.fm5(globalState));
        _1106_v152 = ((p1) ? (_1106_v152) : ((_1108_v154)[_module.__default.safeIndex((_this).f8, new BigNumber((_1108_v154).length))]));
        let _1109_v155;
        let _init28 = function (_1110_i12) {
          return _module.__default.safeDivisionInt(_1110_i12, (_this).f8);
        };
        let _nw184 = Array((new BigNumber(13)).toNumber());
        for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw184.length); _i0_28++) {
          _nw184[_i0_28] = _init28(new BigNumber(_i0_28));
        }
        _1109_v155 = _nw184;
        let _1111_v156;
        _1111_v156 = _dafny.Map.Empty.slice().updateUnsafe(_1109_v155,p2);
        _1111_v156 = (_1111_v156).update(_1109_v155, ((!(p3)) ? ((_this).f8) : (p0)));
        r1 = p0;
        if (p3) {
          let _1112_v157;
          let _nw185 = new _module.C2();
          _nw185.__ctor((_dafny.ZERO).minus(new BigNumber((_1108_v154).length)), (_this).f8);
          _1112_v157 = _nw185;
          (globalState).f3 = p1;
          _1109_v155 = _1109_v155;
          let _1113_v158;
          _1113_v158 = _dafny.Seq.of(p1);
          let _1114_v159;
          _1114_v159 = _dafny.Seq.of(_1113_v158);
          let _1115_v160;
          _1115_v160 = _dafny.Seq.of((_this).f8, p0, new BigNumber(((_1114_v159)[_module.__default.safeIndex(p0, new BigNumber((_1114_v159).length))]).length));
          let _1116_v161;
          _1116_v161 = _dafny.MultiSet.fromElements(p0, (_this).f8, (_dafny.ZERO).minus(p0));
          let _index204 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_1109_v155).length));
          (_1109_v155)[_index204] = (_dafny.ZERO).minus((new BigNumber((_dafny.Seq.UnicodeFromString("ikupdc")).length)).multipliedBy((((_1116_v161).contains((_dafny.ZERO).minus(p2))) ? ((_1116_v161).get((_dafny.ZERO).minus(p2))) : (new BigNumber((_1115_v160).length)))));
          let _1117_v162;
          _1117_v162 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1113_v158).length),_dafny.Seq.Concat(_1115_v160, _1115_v160));
          let _1118_v163;
          _1118_v163 = _dafny.Map.Empty.slice().updateUnsafe(p1,p3);
          let _1119_v164;
          _1119_v164 = _dafny.MultiSet.fromElements((_1114_v159)[_module.__default.safeIndex(new BigNumber((_1118_v163).length), new BigNumber((_1114_v159).length))]);
          let _index205 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_1109_v155).length));
          let _rhs198 = (_this).f7;
          let _rhs199 = (((_1117_v162).contains(p0)) ? ((_1117_v162).get(p0)) : (_1115_v160));
          let _rhs200 = (((_dafny.ZERO).minus(new BigNumber((_1119_v164).cardinality()))).plus(p2)).minus((_this).f7);
          let _lhs147 = _1109_v155;
          let _lhs148 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_1109_v155).length));
          r1 = _rhs198;
          _1115_v160 = _rhs199;
          _lhs147[_lhs148] = _rhs200;
          (globalState).f4 = !(p1);
        } else {
          let _1120_v165;
          _1120_v165 = _dafny.Set.fromElements(p3, p1, p1);
          let _1121_v166;
          _1121_v166 = _dafny.Set.fromElements(p3, p1, (_this).fm1(_1120_v165, p2, new BigNumber(-90), new BigNumber((_1107_v153).length), globalState), !(p1), p3);
          let _1122_v167;
          _1122_v167 = _module.D1.create_DC2();
          let _1123_v168;
          _1123_v168 = _module.D1.create_DC3(_1122_v167);
          let _1124_v169;
          _1124_v169 = _module.D1.create_DC3(_1123_v168);
          let _1125_v170;
          _1125_v170 = _module.D1.create_DC3(_1124_v169);
          let _1126_v171;
          _1126_v171 = _dafny.MultiSet.fromElements(_module.D1.create_DC3(_1124_v169), _1125_v170);
          let _1127_v172;
          let _nw186 = Array((new BigNumber(15)).toNumber());
          _nw186[(_dafny.ZERO).toNumber()] = !(p1);
          _nw186[(_dafny.ONE).toNumber()] = p3;
          _nw186[(new BigNumber(2)).toNumber()] = p1;
          _nw186[(new BigNumber(3)).toNumber()] = p3;
          _nw186[(new BigNumber(4)).toNumber()] = (_this).fm1(_1120_v165, p2, (((_1126_v171).contains(_1125_v170)) ? ((_1126_v171).get(_1125_v170)) : (p0)), (_this).f7, globalState);
          _nw186[(new BigNumber(5)).toNumber()] = p1;
          _nw186[(new BigNumber(6)).toNumber()] = p3;
          _nw186[(new BigNumber(7)).toNumber()] = p1;
          _nw186[(new BigNumber(8)).toNumber()] = p3;
          _nw186[(new BigNumber(9)).toNumber()] = p3;
          _nw186[(new BigNumber(10)).toNumber()] = p3;
          _nw186[(new BigNumber(11)).toNumber()] = p3;
          _nw186[(new BigNumber(12)).toNumber()] = p3;
          _nw186[(new BigNumber(13)).toNumber()] = p1;
          _nw186[(new BigNumber(14)).toNumber()] = p1;
          _1127_v172 = _nw186;
          let _1128_v173;
          _1128_v173 = _module.D6.create_DC15(_1127_v172);
          let _rhs201 = _1128_v173;
          let _rhs202 = ((p1) ? (_1127_v172) : (_1127_v172));
          let _rhs203 = (_this).fm4(p3, p3, globalState);
          let _rhs204 = (_this).f7;
          _1128_v173 = _rhs201;
          _1127_v172 = _rhs202;
          r1 = _rhs203;
          r1 = _rhs204;
          let _index206 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_1127_v172).length));
          (_1127_v172)[_index206] = false;
          let _1129_v174;
          _1129_v174 = p3;
          let _1130_v175;
          let _nw187 = Array((new BigNumber(16)).toNumber());
          _nw187[(_dafny.ZERO).toNumber()] = p1;
          _nw187[(_dafny.ONE).toNumber()] = _1129_v174;
          _nw187[(new BigNumber(2)).toNumber()] = p1;
          _nw187[(new BigNumber(3)).toNumber()] = p1;
          _nw187[(new BigNumber(4)).toNumber()] = _1129_v174;
          _nw187[(new BigNumber(5)).toNumber()] = _1129_v174;
          _nw187[(new BigNumber(6)).toNumber()] = _1129_v174;
          _nw187[(new BigNumber(7)).toNumber()] = _module.__default.fm33(globalState);
          _nw187[(new BigNumber(8)).toNumber()] = _module.__default.fm33(globalState);
          _nw187[(new BigNumber(9)).toNumber()] = _1129_v174;
          _nw187[(new BigNumber(10)).toNumber()] = _1129_v174;
          _nw187[(new BigNumber(11)).toNumber()] = _1129_v174;
          _nw187[(new BigNumber(12)).toNumber()] = _1129_v174;
          _nw187[(new BigNumber(13)).toNumber()] = _1129_v174;
          _nw187[(new BigNumber(14)).toNumber()] = _1129_v174;
          _nw187[(new BigNumber(15)).toNumber()] = _1129_v174;
          _1130_v175 = _nw187;
          let _index207 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_1127_v172).length));
          let _rhs205 = !(p1) || (((_this).fm1(_1121_v166, (_this).f8, (_this).f8, (_this).f8, globalState)) || (p3));
          let _rhs206 = _1130_v175;
          let _lhs149 = _1127_v172;
          let _lhs150 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_1127_v172).length));
          _lhs149[_lhs150] = _rhs205;
          _1130_v175 = _rhs206;
          let _1131_v176;
          _1131_v176 = _dafny.Map.Empty.slice().updateUnsafe(p3,p2);
          let _1132_v177;
          _1132_v177 = _dafny.Seq.of(p3);
          let _index208 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_1127_v172).length));
          let _index209 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_1127_v172).length));
          let _rhs207 = (_module.__default.safeDivisionInt(p2, (_this).f8)).isLessThanOrEqualTo((((_1131_v176).contains(p3)) ? ((_1131_v176).get(p3)) : (p0)));
          let _rhs208 = (_1132_v177)[_module.__default.safeIndex(p2, new BigNumber((_1132_v177).length))];
          let _lhs151 = _1127_v172;
          let _lhs152 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_1127_v172).length));
          let _lhs153 = _1127_v172;
          let _lhs154 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_1127_v172).length));
          _lhs151[_lhs152] = _rhs207;
          _lhs153[_lhs154] = _rhs208;
          let _1133_v178;
          _1133_v178 = _dafny.Seq.of(_dafny.Seq.update(_1108_v154, _module.__default.safeIndex((_this).f8, new BigNumber((_1108_v154).length)), _1106_v152), _1108_v154, _dafny.Seq.UnicodeFromString("bp"), _1107_v153, _1107_v153);
          let _1134_v179;
          _1134_v179 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_1107_v153, (_1133_v178)[_module.__default.safeIndex(p2, new BigNumber((_1133_v178).length))]),_1127_v172);
          let _1135_v180;
          _1135_v180 = _dafny.MultiSet.fromElements(_dafny.Seq.update(_1108_v154, _module.__default.safeIndex(p2, new BigNumber((_1108_v154).length)), _1106_v152));
          let _1136_v181;
          _1136_v181 = _dafny.Set.fromElements((((_1134_v179).contains(_1135_v180)) ? ((_1134_v179).get(_1135_v180)) : (_1127_v172)));
          let _1137_v182;
          _1137_v182 = _module.D5.create_DC11(_1131_v176);
          let _rhs209 = (_1136_v181).Union(_1136_v181);
          let _rhs210 = _1137_v182;
          let _rhs211 = (_1132_v177)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(205)), ((_1138_v152) => function (_1139_i13) {
            return _1138_v152;
          })(_1106_v152)), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(205)), ((_1140_v152) => function (_1141_i13) {
            return _1140_v152;
          })(_1106_v152))).length)), new _dafny.CodePoint('o'.codePointAt(0)))).length), new BigNumber((_1132_v177).length))];
          let _rhs212 = p2;
          let _lhs155 = globalState;
          _1136_v181 = _rhs209;
          _1137_v182 = _rhs210;
          _lhs155.f4 = _rhs211;
          r1 = _rhs212;
          _1132_v177 = _1132_v177;
        }
        r1 = p2;
      } else {
        (globalState).f4 = p3;
        (globalState).f3 = ((new BigNumber(-20)).isLessThan((_this).f8)) === (p1);
        r1 = _module.__default.safeDivisionInt(new BigNumber(193), p0);
        let _1142_v183;
        let _nw188 = Array((new BigNumber(25)).toNumber()).fill(false);
        _1142_v183 = _nw188;
        let _1143_v184;
        _1143_v184 = _module.D6.create_DC15(_1142_v183);
        _1142_v183 = (_1143_v184).dtor_cf23;
        let _1144_v185;
        _1144_v185 = _dafny.Seq.of(p1, p3, true);
        let _1145_v186;
        _1145_v186 = _module.D6.create_DC16((_module.__default.fm35(p2, globalState)).dtor_cf24, _1144_v185);
        let _1146_v187;
        _1146_v187 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(p1),p0);
        let _1147_v188;
        _1147_v188 = _dafny.Set.fromElements(_module.D6.create_DC16(_1146_v187, _1144_v185));
        let _1148_v189;
        _1148_v189 = _dafny.Seq.of(_module.__default.fm34(p1, p3, globalState), (_dafny.Set.fromElements(_1145_v186, _1145_v186)).Intersect(_dafny.Set.fromElements(_1145_v186, _1145_v186)), _1147_v188, _1147_v188, (_1147_v188).Union(_1147_v188));
        _1148_v189 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(705)), ((_1149_v188) => function (_1150_i14) {
          return _1149_v188;
        })(_1147_v188));
      }
      let _1151_v190;
      let _nw189 = Array((new BigNumber(24)).toNumber()).fill(false);
      _1151_v190 = _nw189;
      let _index210 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_1151_v190).length));
      (_1151_v190)[_index210] = p3;
      let _index211 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_1151_v190).length));
      (_1151_v190)[_index211] = p3;
      r1 = p2;
      let _1152_v191;
      _1152_v191 = _dafny.Seq.of(p3);
      r0 = _1152_v191;
      let _1153_v192;
      _1153_v192 = _module.D5.create_DC11((_dafny.Map.Empty.slice().updateUnsafe((_1151_v190)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_1151_v190).length))],p2)).update(p1, (_this).f8));
      let _pat_let_tv15 = p0;
      let _pat_let_tv16 = _1151_v190;
      let _pat_let_tv17 = _1151_v190;
      let _pat_let_tv18 = p0;
      let _pat_let_tv19 = p3;
      r1 = (_dafny.ZERO).minus(function (_source23) {
        if (_source23.is_DC12) {
          let _1154___mcc_h10 = (_source23).cf19;
          let _1155_cf19 = _1154___mcc_h10;
          return (_dafny.ZERO).minus(_pat_let_tv15);
        } else if (_source23.is_DC13) {
          let _1156___mcc_h11 = (_source23).cf20;
          let _1157___mcc_h12 = (_source23).cf21;
          let _1158_cf21 = _1157___mcc_h12;
          let _1159_cf20 = _1156___mcc_h11;
          return new BigNumber(862);
        } else if (_source23.is_DC11) {
          let _1160___mcc_h13 = (_source23).cf18;
          let _1161_cf18 = _1160___mcc_h13;
          return (_this).f8;
        } else {
          let _1162___mcc_h14 = (_source23).cf22;
          let _1163_cf22 = _1162___mcc_h14;
          return _module.__default.safeModuloInt((_module.D2.create_DC5((_module.D3.create_DC7((_pat_let_tv17)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_pat_let_tv16).length))])).dtor_cf10, (_this).f7, (_this).f8, _pat_let_tv18, (_pat_let_tv19))).dtor_cf7, new BigNumber(-174));
        }
      }(_1153_v192));
      r2 = !(p3);
      return [r0, r1, r2];
    }
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let _1164_v0;
      let _nw190 = new _module.C2();
      _nw190.__ctor((_this).f8, new BigNumber(687));
      _1164_v0 = _nw190;
      let _1165_v1;
      _1165_v1 = _dafny.Set.fromElements(_1164_v0, _1164_v0, _1164_v0, _1164_v0, _1164_v0);
      let _1166_v2;
      _1166_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,(_1165_v1).Intersect(_1165_v1));
      _1165_v1 = (((_1166_v2).contains(p1)) ? ((_1166_v2).get(p1)) : (_1165_v1));
      let _1167_v3;
      _1167_v3 = _dafny.Seq.of((_this).f8);
      r0 = (_1167_v3)[_module.__default.safeIndex(p3, new BigNumber((_1167_v3).length))];
      let _1168_v4;
      let _nw191 = new _module.C0();
      _nw191.__ctor(p2, (_this).f8, (_module.__default.fm20(p0, p0, p0, globalState)).plus((_this).f8));
      _1168_v4 = _nw191;
      _1168_v4 = _1168_v4;
      let _1169_v5;
      _1169_v5 = _module.D4.create_DC10((_dafny.ZERO).minus((_this).f7), (p3).isLessThan((_this).f8));
      let _1170_v6;
      let _init29 = ((_1171_p0) => function (_1172_i0) {
        return (_dafny.MultiSet.fromElements(_1171_p0)).Difference(_dafny.MultiSet.fromElements(_1171_p0, _1171_p0, _1171_p0, _1171_p0, _1171_p0));
      })(p0);
      let _nw192 = Array((new BigNumber(11)).toNumber());
      for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw192.length); _i0_29++) {
        _nw192[_i0_29] = _init29(new BigNumber(_i0_29));
      }
      _1170_v6 = _nw192;
      let _rhs213 = _1169_v5;
      let _rhs214 = _1170_v6;
      _1169_v5 = _rhs213;
      _1170_v6 = _rhs214;
      let _hi8 = new BigNumber(-776);
      for (let _1173_i1 = new BigNumber(307); _1173_i1.isLessThan(_hi8); _1173_i1 = _1173_i1.plus(_dafny.ONE)) {
        let _1174_v7;
        _1174_v7 = _module.D7.create_DC19(p1, p0, (_1168_v4).f7);
        let _1175_v8;
        let _init30 = ((_1176_p0) => function (_1177_i2) {
          return _1176_p0;
        })(p0);
        let _nw193 = Array((new BigNumber(2)).toNumber());
        for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw193.length); _i0_30++) {
          _nw193[_i0_30] = _init30(new BigNumber(_i0_30));
        }
        _1175_v8 = _nw193;
        let _1178_v9;
        _1178_v9 = _dafny.MultiSet.fromElements(_1175_v8, _1175_v8, _1175_v8, _1175_v8);
        let _1179_v10;
        _1179_v10 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1168_v4,p0)).length));
        let _1180_v11;
        let _nw194 = new _module.C4();
        _nw194.__ctor((_1179_v10).IsProperSubsetOf(_dafny.MultiSet.fromElements((_1174_v7).dtor_cf29, new BigNumber((_1178_v9).cardinality()))), (_1168_v4).f8, (_1168_v4).f8);
        _1180_v11 = _nw194;
        let _1181_v12;
        _1181_v12 = new _dafny.CodePoint('y'.codePointAt(0));
        _1181_v12 = _1181_v12;
        let _1182_v13;
        _1182_v13 = _dafny.Seq.of(p0, p0);
        let _1183_v14;
        _1183_v14 = _dafny.Seq.of(_1182_v13);
        let _1184_v15;
        let _nw195 = new _module.C3();
        _nw195.__ctor((_1180_v11).f11, (_1168_v4).f7, (_this).f7);
        _1184_v15 = _nw195;
        let _1185_v16;
        _1185_v16 = _dafny.Set.fromElements(_1184_v15);
        let _1186_v17;
        let _nw196 = new _module.C3();
        _nw196.__ctor(!(_dafny.areEqual(_1182_v13, (_1183_v14)[_module.__default.safeIndex(p1, new BigNumber((_1183_v14).length))])), new BigNumber(17), (new BigNumber((_1185_v16).length)).plus(p1));
        _1186_v17 = _nw196;
        let _1187_v18;
        _1187_v18 = _dafny.Seq.UnicodeFromString("eu");
        let _1188_v19;
        _1188_v19 = _module.D8.create_DC21(_1184_v15);
        let _rhs215 = (_1188_v19).dtor_cf34;
        let _rhs216 = (_1168_v4).f7;
        let _rhs217 = _dafny.Seq.update((_1186_v17).fm2((_1186_v17).f12, globalState), _module.__default.safeIndex((_1168_v4).f8, new BigNumber(((_1186_v17).fm2((_1186_v17).f12, globalState)).length)), (((_1184_v15).f12) ? (_1181_v12) : (_1181_v12)));
        _1186_v17 = _rhs215;
        r0 = _rhs216;
        _1187_v18 = _rhs217;
        r2 = (_1184_v15).f12;
      }
      let _1189_v20;
      _1189_v20 = new _dafny.CodePoint('a'.codePointAt(0));
      let _1190_v21;
      _1190_v21 = _dafny.Set.fromElements(false);
      let _1191_v22;
      _1191_v22 = _dafny.Map.Empty.slice().updateUnsafe(!(true),_1190_v21);
      let _1192_v23;
      _1192_v23 = _dafny.Map.Empty.slice().updateUnsafe((_1168_v4).f7,(((_1191_v22).contains(true)) ? ((_1191_v22).get(true)) : (_1190_v21)));
      let _1193_v24;
      _1193_v24 = _dafny.Seq.of(p0, p0);
      let _1194_v25;
      _1194_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1189_v20,(((_1192_v23).contains(new BigNumber((_1193_v24).length))) ? ((_1192_v23).get(new BigNumber((_1193_v24).length))) : (_dafny.Set.fromElements(!(p0)))));
      _1194_v25 = (_1194_v25).update(_1189_v20, _1190_v21);
      r0 = new BigNumber(872);
      r1 = (_1168_v4).f7;
      r2 = p0;
      return [r0, r1, r2];
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
